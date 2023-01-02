import { createI18n } from 'vue-i18n'
import messages from './i18n.js'

const i18n = createI18n({
  locale: localStorage.lang || 'cn',
  messages,
  datetimeFormats: {
    'cn': {
      s: {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
      },
      l: {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false,
      },
      short: {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
      },
      long: {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        weekday: 'short',
        hour: 'numeric',
        minute: 'numeric',
      },
    },
    'en': {
      s: {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
      },
      l: {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false,
      },
      short: {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
      },
      long: {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        weekday: 'short',
        hour: 'numeric',
        minute: 'numeric',
      },
    },
  },

  modifiers: {},
})

export default i18n
