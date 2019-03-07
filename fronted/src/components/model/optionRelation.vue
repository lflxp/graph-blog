<template>
  <div>
    <a-row>
      <a-col :span="3">
        <a-menu
          style="width: 100%;"
          :defaultSelectedKeys="['1']"
          :defaultOpenKeys="['2']"
          :selectedKeys="selects"
          @click="handleClick"
          @dbclick="handleClick2"
          mode="inline"
          theme="dark"
          :inlineCollapsed="collapsed"
        >
          <template v-for="item in list">
            <a-menu-item v-if="!item.children" :key="item.key">
              <a-icon type="pie-chart" />
              <span>{{ item.title }}</span>
            </a-menu-item>
            <sub-menu v-else :menu-info="item" :key="item.key" :showadd="false" :send="edit" :show="showkeypath" />
          </template>
        </a-menu>
      </a-col>
      <a-col :span="21" style="padding: 0 10px;">
        <d3-manul :nodes="nodes" :edges="edges" v-if="resetd3"></d3-manul>
      </a-col>
    </a-row>
    <a-drawer
      title="添加类型关系"
      :width="720"
      placement="right"
      :closable="false"
      @close="onClose"
      :visible="visible"
      :wrapStyle="{height: 'calc(100% - 108px)',overflow: 'auto',paddingBottom: '108px'}">
      <a-form :form="form" layout="vertical" hideRequiredMark>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="类型组名">
              <a-input
                v-model="menuConfig.group"
                :disabled="true"/>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="类型名称">
              <a-input
                v-model="menuConfig.name"
                :disabled="true"/>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-tabs defaultActiveKey="1" @change="callback">
            <a-tab-pane tab="添加上层类型关系" key="1">
              <a-select style="width: 100%;" @change="handleFromChange" placeholder="修改上层类型关系请选择">
                <a-select-option
                  v-for="(item,index) in menuConfig.relation.from"
                  :key="item.relation"
                  :value="index">
                  {{ item.relation }} {{ item.item }}
                </a-select-option>
              </a-select>
              <a-timeline style="padding: 20px 0 0 0;">
                <a-timeline-item>
                  <a-select
                    mode="multiple"
                    v-model="from.item"
                    placeholder="请选择类型">
                    <a-select-opt-group
                      v-for="item in list"
                      :label="item.key"
                      :key="item.key">
                      <a-select-option
                        v-for="ch in item.children"
                        :key="ch.key"
                        :value="item.key+','+ch.key">
                        {{ ch.key }}
                      </a-select-option>
                    </a-select-opt-group>
                  </a-select>
                </a-timeline-item>
                <a-timeline-item color="red">
                  <a-icon slot="dot" type="arrow-down" style="fontSize: '16px'" />
                  <!-- <a-select style="width: 100%;" @change="handleSelectChange" placeholder="请选择关系类型">
                    <a-select-option value="虚拟化">虚拟化</a-select-option>
                    <a-select-option value="管理">管理</a-select-option>
                    <a-select-option value="运行在">运行在</a-select-option>
                    <a-select-option value="连接">连接</a-select-option>
                    <a-select-option value="提供">提供</a-select-option>
                    <a-select-option value="使用">使用</a-select-option>
                    <a-select-option value="依赖">依赖</a-select-option>
                    <a-select-option value="包含">包含</a-select-option>
                    <a-select-option value="组成">组成</a-select-option>
                  </a-select> -->
                  <a-auto-complete
                    :dataSource="dataSource"
                    v-model="from.relation"
                    style="width: 100%;"
                    placeholder="请选择关系类型"
                    :filterOption="filterOption"/>
                </a-timeline-item>
                <a-timeline-item color="green">
                  <a-icon slot="dot" type="arrow-down" style="fontSize: '16px'" />
                  <a-select style="width: 100%;" @change="handleSelectChange" placeholder="请选择限制条件" v-model="from.limit">
                    <a-select-option value="1:1">1:1</a-select-option>
                    <a-select-option value="1:N">1:N</a-select-option>
                    <a-select-option value="N:1">N:1</a-select-option>
                    <a-select-option value="N:N">N:N</a-select-option>
                  </a-select>
                </a-timeline-item>
                <a-timeline-item>
                  <a-icon slot="dot" type="check-square" style="fontSize: '16px'" />
                  <a-button type="dashed" style="width: 100%;">{{ menuConfig.name }}</a-button>
                </a-timeline-item>
                <a-button type="danger" @click="deletefrom" v-if="showdeletefrom">删除</a-button>
              </a-timeline>
            </a-tab-pane>
            <a-tab-pane tab="添加下层类型关系" key="2" forceRender>
              <a-select style="width: 100%;" @change="handleToChange" placeholder="修改下层类型关系请选择">
                <a-select-option
                  v-for="(item,index) in menuConfig.relation.to"
                  :key="item.relation"
                  :value="index">
                  {{ item.relation }} {{ item.item }}
                </a-select-option>
              </a-select>
              <a-timeline style="padding: 20px 0 0 0;">
                <a-timeline-item>
                  <a-button type="dashed" style="width: 100%;">{{ menuConfig.name }}</a-button>
                </a-timeline-item>
                <a-timeline-item color="red">
                  <a-icon slot="dot" type="arrow-down" style="fontSize: '16px'" />
                  <!-- <a-select style="width: 100%;" @change="handleSelectChange" placeholder="请选择关系类型">
                    <a-select-option value="虚拟化">虚拟化</a-select-option>
                    <a-select-option value="管理">管理</a-select-option>
                    <a-select-option value="运行在">运行在</a-select-option>
                    <a-select-option value="连接">连接</a-select-option>
                    <a-select-option value="提供">提供</a-select-option>
                    <a-select-option value="使用">使用</a-select-option>
                    <a-select-option value="依赖">依赖</a-select-option>
                    <a-select-option value="包含">包含</a-select-option>
                    <a-select-option value="组成">组成</a-select-option>
                  </a-select> -->
                  <a-auto-complete
                    :dataSource="dataSource"
                    style="width: 100%;"
                    v-model="to.relation"
                    placeholder="请选择关系类型"
                    :filterOption="filterOption"/>
                </a-timeline-item>
                <a-timeline-item color="green">
                  <a-icon slot="dot" type="arrow-down" style="fontSize: '16px'" />
                  <a-select style="width: 100%;" @change="handleSelectChange" placeholder="请选择限制条件" v-model="to.limit">
                    <a-select-option value="1:1">1:1</a-select-option>
                    <a-select-option value="1:N">1:N</a-select-option>
                    <a-select-option value="N:1">N:1</a-select-option>
                    <a-select-option value="N:N">N:N</a-select-option>
                  </a-select>
                </a-timeline-item>
                <a-timeline-item>
                  <a-icon slot="dot" type="check-square" style="fontSize: '16px'" />
                  <a-select
                    mode="multiple"
                    v-model="to.item"
                    placeholder="请选择类型">
                    <a-select-opt-group
                      v-for="item in list"
                      :label="item.key"
                      :key="item.key">
                      <a-select-option
                        v-for="ch in item.children"
                        :key="ch.key"
                        :value="item.key+','+ch.key">
                        {{ ch.key }}
                      </a-select-option>
                    </a-select-opt-group>
                  </a-select>
                </a-timeline-item>
                <a-button type="danger" @click="deleteto" v-if="showdeleteto">删除</a-button>
              </a-timeline>
            </a-tab-pane>
          </a-tabs>
        </a-row>
      </a-form>
      <div
        :style="{
          position: 'absolute',
          left: 0,
          bottom: 0,
          width: '100%',
          borderTop: '1px solid #e9e9e9',
          padding: '10px 16px',
          background: '#fff',
          textAlign: 'right',
        }"
      >
        <a-button
          :style="{marginRight: '8px'}"
          @click="onClose"
        >
          取消
        </a-button>
        <a-button @click="submit" type="primary">提交</a-button>
      </div>
    </a-drawer>
  </div>
</template>

<script>
import SubMenu from './SubMenu1'
import Utils from '@/utils/warehouse'
import D3Manul from '@/components/D3Pic/manul'
import icons from './icons'

// const nodesinit = [
//   { 'group': 'a', 'id': 50, 'labels': ['Person'], 'props': { 'born': 1956, 'name': 'Tom Hanks' }, 'type': 'node' },
//   { 'group': 'b', 'id': 163, 'labels': ['Movie'], 'props': { 'released': 1995, 'tagline': 'Houston, we have a problem.', 'title': 'Apollo 13' }, 'type': 'node' },
//   { 'group': 'm', 'id': 134, 'labels': ['Person'], 'props': { 'born': 1954, 'name': 'Ron Howard' }, 'type': 'node' },
//   { 'group': 'a', 'id': 50, 'labels': ['Person'], 'props': { 'born': 1956, 'name': 'Tom Hanks' }, 'type': 'node' },
//   { 'group': 'b', 'id': 181, 'labels': ['Movie'], 'props': { 'released': 1992, 'tagline': 'Once in a lifetime you get a chance to do something different.', 'title': 'A League of Their Own' }, 'type': 'node' },
//   { 'group': 'm', 'id': 185, 'labels': ['Person'], 'props': { 'born': 1943, 'name': 'Penny Marshall' }, 'type': 'node' },
//   { 'group': 'a', 'id': 50, 'labels': ['Person'], 'props': { 'born': 1956, 'name': 'Tom Hanks' }, 'type': 'node' },
//   { 'group': 'b', 'id': 57, 'labels': ['Movie'], 'props': { 'released': 1990, 'tagline': 'A story of love, lava and burning desire.', 'title': 'Joe Versus the Volcano' }, 'type': 'node' }]

// const edgesinit = [
//   { 'name': 'r', 'props': { 'roles': ['Jim Lovell'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 1, 'value': 1 },
//   { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 6, 'target': 1, 'value': 1 },
//   { 'name': 'r', 'props': { 'roles': ['Jimmy Dugan'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 4, 'value': 1 },
//   { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 5, 'target': 4, 'value': 1 },
//   { 'name': 'r', 'props': { 'roles': ['Joe Banks'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 2, 'value': 1 },
//   { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 4, 'target': 7, 'value': 3 },
//   { 'name': 'r', 'props': { 'roles': ['Mr. White'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 4, 'value': 1 },
//   { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 3, 'target': 6, 'value': 1 }]
const Base64 = require('js-base64').Base64
export default {
  components: {
    D3Manul,
    Utils,
    'sub-menu': SubMenu
  },
  watch: {
    keypath () {
      if (this.keypath.length > 0) {
        this.showkeypath = true
      } else {
        this.showkeypath = false
      }
    }
  },
  data () {
    return {
      resetd3: true,
      reverse: false,
      nodes: [],
      edges: [],
      info: 'Tom',
      node: 'Person',
      num1: 25,
      groupname: '',
      visiblemodel: false,
      visible: false,
      keypath: [],
      showkeypath: false,
      current: '1',
      theme: 'dark',
      selects: [],
      currentdata: new Map(),
      collapsed: false,
      list: [],
      icons: icons,
      form: this.$form.createForm(this),
      dataSource: ['虚拟化', '管理', '运行在', '连接', '提供', '使用', '依赖', '包含', '组成'],
      from: {
        item: [],
        relation: '',
        limit: ''
      },
      fromindex: 0,
      showdeletefrom: false,
      to: {
        item: [],
        relation: '',
        limit: ''
      },
      toindex: 0,
      showdeleteto: false,
      menuConfig: {
        group: '',
        level: 1,
        name: '',
        node: '',
        icon: '',
        columns: [],
        template: [],
        created: '',
        isgroup: false,
        isslave: true,
        relation: {
          from: [],
          to: []
        }
      }
    }
  },
  mounted () {
    this.refresh()
  },
  methods: {
    iconclick (item) {
      console.log('iconclick', item)
    },
    handleFromChange (key) {
      // console.log('handleFromChange', key)
      this.showdeletefrom = true
      this.from = this.menuConfig.relation.from[key]
      this.fromindex = key
    },
    deleteto () {
      if (this.toindex > -1) {
        this.menuConfig.relation.to.splice(this.toindex, 1)
        // console.log('deleteto', this.menuConfig)
        if (this.menuConfig.name.length > 0) {
          var tmp = {
            'key': '/ares/cmdb/model/' + this.menuConfig.group + '/' + this.menuConfig.name,
            'ttl': 1,
            'value': JSON.stringify(this.menuConfig)
          }
          Utils.setInfo(tmp).then(resp => {
            this.refresh()
            this.refreshto()
            this.$message.success(resp.data.status)
          })
        } else {
          this.$message.error('对象名称为空，请选择！')
        }
      } else {
        this.$message.error('索引error' + this.toindex)
      }
    },
    deletefrom () {
      if (this.fromindex > -1) {
        this.menuConfig.relation.from.splice(this.fromindex, 1)
        // console.log('deletefrom', this.menuConfig)
        if (this.menuConfig.name.length > 0) {
          var tmp = {
            'key': '/ares/cmdb/model/' + this.menuConfig.group + '/' + this.menuConfig.name,
            'ttl': 1,
            'value': JSON.stringify(this.menuConfig)
          }
          Utils.setInfo(tmp).then(resp => {
            this.refresh()
            this.refreshfrom()
            this.$message.success(resp.data.status)
          })
        } else {
          this.$message.error('对象名称为空，请选择！')
        }
      } else {
        this.$message.error('索引error' + this.fromindex)
      }
    },
    handleToChange (key) {
      // console.log('handleToChange', key)
      this.showdeleteto = true
      this.to = this.menuConfig.relation.to[key]
      this.toindex = key
    },
    filterOption (input, option) {
      return option.componentOptions.children[0].text.toUpperCase().indexOf(input.toUpperCase()) >= 0
    },
    handletimeClick () {
      this.reverse = !this.reverse
    },
    callback (key) {
      console.log(key)
    },
    handleSelectChange (i) {
      console.log(i)
    },
    onChange (i) {
      console.log(i)
    },
    submit () {
      if (this.menuConfig.name.length > 0) {
        this.menuConfig.created = (new Date()).toDateString() + ' ' + (new Date()).toTimeString()
        // console.log('submit from', this.from)
        if (this.from.relation.length > 0 && this.from.item.length > 0 && this.from.limit.length > 0) {
          this.menuConfig.relation.from.push(this.from)
        }
        // console.log('submit to', this.to)
        if (this.to.relation.length > 0 && this.to.item.length > 0 && this.to.limit.length > 0) {
          this.menuConfig.relation.to.push(this.to)
        }
        var tmp = {
          'key': '/ares/cmdb/model/' + this.menuConfig.group + '/' + this.menuConfig.name,
          'ttl': 1,
          'value': JSON.stringify(this.menuConfig)
        }
        // console.log(tmp)
        Utils.setInfo(tmp).then(resp => {
          // console.log(resp)
          this.refresh()
          this.$message.success(resp.data.status)
        })
        this.visiblemodel = false
      } else {
        this.$message.error('内容为空，请输入组名')
      }
    },
    cancel () {
      this.visible = false
    },
    refresh () {
    //   this.$router.go(0)
      var se = {
        'key': '/ares/cmdb/model',
        'recursive': true
      }
      Utils.getInfo(se).then(data => {
        // console.log(data)
        //   var tmp_list = new Map()
        // var decodedData = window.atob(encodedData);
        if (data.data.data.kvs !== undefined) {
          this.list = []
          var tmpData = new Map()
          data.data.data.kvs.forEach((a) => {
            var tmp = JSON.parse(Base64.decode(a.value))
            this.currentdata.set(tmp.name, tmp)
            if (tmp.isgroup === true && tmp.isslave === false) {
              tmpData.set(tmp.name, { key: tmp.name, title: tmp.name, children: [] })
            }
          })
          data.data.data.kvs.forEach((b, index) => {
            var tmp1 = JSON.parse(Base64.decode(b.value))
            if (tmp1.isgroup === false && tmp1.isslave === true) {
              if (tmpData.has(tmp1.group)) {
                var t1 = tmpData.get(tmp1.group)
                if (t1.children !== undefined && t1.children.length > 0) {
                  t1.children.push({ key: tmp1.name, title: tmp1.name, icon: tmp1.icon, group: tmp1.group, level: index })
                  tmpData.set(tmp1.group, t1)
                } else {
                  t1.children = [{ key: tmp1.name, title: tmp1.name, icon: tmp1.icon, group: tmp1.group, level: index }]
                  tmpData.set(tmp1.group, t1)
                }
              }
            }
          })
          tmpData.forEach((value, key, map) => {
            this.list.push(value)
          })
        }
        // console.log('finished', this.list)
      })
    },
    showDrawer () {
      this.visible = true
    },
    onClose () {
      this.visible = false
      // 强制用户点击后才能编辑
      this.keypath = []
    },
    changeTheme (checked) {
      this.theme = checked ? 'dark' : 'light'
    },
    onSearch (value) {
      console.log(value)
    },
    handleClick (e) {
      // console.log('handleClick', e, this.currentdata)
      //   this.menuConfig.name = e.keyPath[0]
      //   this.menuConfig.group = e.keyPath[1]
      this.menuConfig = this.currentdata.get(e.keyPath[0])
      console.log(this.menuConfig)
      this.keypath = e.keyPath
      this.createD3data()
      this.refreshd3()
      // this.handleClick2(e)
    //   this.$refs.table.refresh(true)
    },
    refreshd3 () {
      this.resetd3 = false
      this.$nextTick(() => {
        this.resetd3 = true
      })
    },
    createD3data () {
      this.nodes = []
      this.edges = []
      if (this.nodes.length === 0) {
        this.nodes.push({
          group: 'init',
          labels: [this.menuConfig.node],
          props: {
            name: this.menuConfig.name
          },
          type: 'node'
        })
      }
      // from parse
      this.menuConfig.relation.from.forEach((f) => {
        // console.log('f', f)
        f.item.forEach((ff) => {
          this.nodes.push({
            group: 'from',
            labels: [ff],
            props: {
              name: ff.split(',')[1]
            },
            type: 'node'
          })
          this.edges.push({
            name: 'from',
            relation: f.relation,
            limit: f.limit,
            source: this.nodes.length - 1,
            target: 0
          })
        })
      })
      // to parse
      this.menuConfig.relation.to.forEach((f) => {
        // console.log('f', f)
        f.item.forEach((ff) => {
          this.nodes.push({
            group: 'from',
            labels: [ff],
            props: {
              name: ff.split(',')[1]
            },
            type: 'node'
          })
          this.edges.push({
            name: 'from',
            relation: f.relation,
            limit: f.limit,
            source: 0,
            target: this.nodes.length - 1
          })
        })
      })
      // console.log('createD3data', this.nodes, this.edges)
    },
    refreshfrom () {
      this.from.item = []
      this.from.relation = ''
      this.from.limit = ''
    },
    refreshto () {
      this.to.item = []
      this.to.relation = ''
      this.to.limit = ''
    },
    edit (name) {
      // console.log('edit', name)
      // console.log('menuConfig', this.menuConfig)
      // console.log('keypath', this.keypath)
      this.menuConfig = this.currentdata.get(name.title)
      this.refreshfrom()
      this.refreshto()
      this.menuConfig.group = name.group
      this.showDrawer()
    //   this.$refs.table.refresh(true)
    },
    handleClick2 (e) {
      // this.refresh()
      //   console.log('handleClick222', e, this.keyPath)
      //   console.log(e.target.outerText)
      // menu 默认为2级目录 超过数据会出错
      // menu 默认也不能为1级信息
      // console.log(e)
      this.menuConfig = this.currentdata.get(e.target.outerText)
      this.refreshfrom()
      if (this.keypath.length === 2) {
        // 点击非叶子节点
        // console.log('1', e.target.outerText, this.keypath)
        if (e.target.outerText === this.keypath[0] && e.target.outerText !== this.keypath[1]) {
          this.menuConfig.group = this.keypath[1]
        } else {
          this.menuConfig.group = e.target.outerText
        }
      } else if (this.keypath.length === 1) {
        if (e.target.outerText === this.keypath[0]) {
        //   console.log('1.1', e.target.outerText, this.keypath)
          this.menuConfig.group = e.target.outerText
        } else {
        //   console.log('2.1', e.target.outerText, this.keypath)
          this.menuConfig.group = this.keypath[0]
        }
      } else {
        // 用户未点击子菜单的情况下
        // console.log('3', e.target.outerText, this.keypath)
        this.menuConfig.group = e.target.outerText
      }
      this.showDrawer()
    //   this.$refs.table.refresh(true)
    }
  }
}
</script>
