import { ref } from 'vue'

const defaultOptions = {
  autoReconnect: true,
  heartbeat: true,
  onConnected: (ws) => {},
  onDisconnected: (ws, event) => { },
  onError: (ws, event) => { },
  onFailed: () => { },
  onMessage: (ws, event) => { },
}

const DEFAULT_PING_MESSAGE = 'ping'

export default function useWebsocket(url, options = defaultOptions) {
  let explicitlyClosed = false
  let retried = 0
  const status = ref('CLOSED')
  const wsRef = ref()
  let bufferedData = []
  let heartbeatPause
  let heartbeatResume
  let pongTimeoutWait

  const data = ref(null)

  options = { ...defaultOptions, ...options }
  const {
    onConnected,
    onDisconnected,
    onError,
    onMessage,
    onFailed,
  } = options

  // heartbeat feature
  if (options.heartbeat) {
    const isActive = ref(false) // heartbeat active status
    const message = DEFAULT_PING_MESSAGE
    const interval = 5000 // time interval for heartbeat
    const pongTimeout = 10000 // time to stop waiting pong response

    let _timer = null

    const ping = () => {
      send(message, false)
      if (pongTimeoutWait !== undefined) { return }
      pongTimeoutWait = setTimeout(() => {
        // auto-reconnect will be trigger with ws.onclose()
        close()
        explicitlyClosed = false
      }, pongTimeout)
    }

    heartbeatPause = () => {
      isActive.value = false
      clean()
    }

    heartbeatResume = () => {
      isActive.value = true
      clean()
      _timer = setInterval(ping, interval)
    }

    const clean = () => {
      if (_timer) {
        clearInterval(_timer)
        _timer = null
      }
    }
  }

  const _sendBuffer = () => {
    if (bufferedData.length && wsRef.value && status.value === 'OPEN') {
      for (const buffer of bufferedData) { wsRef.value.send(buffer) }
      bufferedData = []
    }
  }

  const resetHeartbeat = () => {
    clearTimeout(pongTimeoutWait)
    pongTimeoutWait = undefined
  }

  const send = (data, useBuffer = true) => {
    if (!wsRef.value || status.value !== 'OPEN') {
      if (useBuffer) { bufferedData.push(data) }
      return false
    }
    _sendBuffer()
    wsRef.value.send(data)
    return true
  }

  const _init = () => {
    if (explicitlyClosed) { return }
    const ws = new WebSocket(url)
    wsRef.value = ws
    status.value = 'CONNECTING'

    ws.onopen = () => {
      status.value = 'OPEN'
      onConnected?.(ws)
      heartbeatResume?.()
      _sendBuffer()
    }

    ws.onclose = (ev) => {
      console.log('---ws.onclose----')
      status.value = 'CLOSED'
      wsRef.value = undefined
      onDisconnected?.(ws, ev)

      if (!explicitlyClosed && options.autoReconnect) {
        console.log('---start retry...----')
        const retries = 5 // infinite retry reconnect
        const delay = 60000
        retried += 1
        if (typeof retries === 'number' && (retries < 0 || retried < retries)) {
          setTimeout(_init, delay)
        } else {
          onFailed?.()
        }
      }
    }

    ws.onerror = (e) => {
      onError?.(ws, e)
    }

    ws.onmessage = (e) => {
      if (options.heartbeat) {
        resetHeartbeat()
        if (e.data === DEFAULT_PING_MESSAGE) {
          return
        }
      }

      data.value = e.data
      onMessage?.(ws, e)
    }
  }

  const open = () => {
    close()
    explicitlyClosed = false
    retried = 0
    _init()
  }

  // Status code 1000 -> Normal Closure https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent/code
  const close = (code = 1000, reason) => {
    if (!wsRef.value) {
      return
    }
    explicitlyClosed = true
    resetHeartbeat()
    heartbeatPause?.()
    wsRef.value.close(code, reason)
  }

  return {
    data,
    status,
    close,
    send,
    open,
  }
}
