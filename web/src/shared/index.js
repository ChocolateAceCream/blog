import MyForm from './components/MyForm'
import MyTable from './components/MyTable'
import Pagination from './components/Pagination'
import Modal from './components/Modal'
import Permission from './components/Permission'
import SvgIcon from './components/SvgIcon'

export default {
  install(app) {
    app.component('MyTable', MyTable)
    app.component('MyForm', MyForm)
    app.component('Pagination', Pagination)
    app.component('Permission', Permission)
    app.component('SvgIcon', SvgIcon)
    app.component('Modal', Modal)
  }
}
