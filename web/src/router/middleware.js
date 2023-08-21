import NProgress from 'nprogress'

export const routeMiddleware = (routes = []) => {
  const wrappedRouter = routes.map(route => {
    return routeFormation(route)
  })
  return wrappedRouter
}

function routeFormation(route) {
  const result = addNProgress(route)
  if (route.children?.length > 0) {
    result['children'] = route.children.map(childRoute => {
      return routeFormation(childRoute)
    })
  }
  return result
}


// apply NProgress to async loaded routes
function addNProgress(options = {}) {
  const { component, ...finalOptions } = options
  if (typeof options['component'] === 'function') {
    finalOptions['component'] = () => {
      const result = options['component']()
      if (typeof result === 'object' && typeof result.then === 'function') {
        NProgress.start()
        result.finally(() => {
          NProgress.done()
        })
      }
      return result
    }
  } else {
    finalOptions['component'] = component
  }
  return finalOptions
}
