import MyTable from './MyTable'
import Pagination from './Pagination'
import Modal from './Modal'
import Permission from './Permission'
import SvgIcon from './SvgIcon'

export default {
  install(app) {
    app.use(MyTable)
    app.use(Pagination)
    app.use(Modal)
    app.use(Permission)
    app.use(SvgIcon)
  }
}
