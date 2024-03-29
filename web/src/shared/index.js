import MySelect from './components/MySelect'
import MyForm from './components/MyForm'
import MyTree from './components/MyTree'
import MyTable from './components/MyTable'
import Pagination from './components/Pagination'
import Modal from './components/Modal'
import Permission from './components/Permission'
import SvgIcon from './components/SvgIcon'
import MyEditor from './components/MyEditor'
import MySearchBar from './components/MySearchBar'
import VDebounce from './directive/debounce'
import VKeyboard from './directive/keyboard'

export default {
  install(app) {
    app.component('MyTable', MyTable)
    app.component('MyTree', MyTree)
    app.component('MyForm', MyForm)
    app.component('MySelect', MySelect)
    app.component('Pagination', Pagination)
    app.component('Permission', Permission)
    app.component('SvgIcon', SvgIcon)
    app.component('Modal', Modal)
    app.component('MyEditor', MyEditor)
    app.component('MySearchBar', MySearchBar)

    // directives
    app.directive('debounce', VDebounce)
    app.directive('keyboard', VKeyboard)
  }
}
