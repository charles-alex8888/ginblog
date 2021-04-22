import Vue from 'vue'
import { Button, FormModel, Input, Icon, message, Layout, Menu, Card, Table, Row, Col, ConfigProvider, Modal, Select } from 'ant-design-vue'

message.config({
  duration: 2,
  top: '60px',
  maxCount: 3
})
Vue.prototype.$message = message
Vue.prototype.$confirm = Modal.confirm

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Table)
Vue.use(Col)
Vue.use(Row)
Vue.use(ConfigProvider)
Vue.use(Modal)
Vue.use(Select)
