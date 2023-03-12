import MyTable from './components/MyTable'
import Pagination from './components/Pagination'
import Modal from './components/Modal'
import Permission from './components/Permission'
import SvgIcon from './components/SvgIcon'

export default {
  install(app) {
    app.use(MyTable)
    app.use(Pagination)
    app.use(Modal)
    app.use(Permission)
    app.component('SvgIcon', SvgIcon)
  }
}
