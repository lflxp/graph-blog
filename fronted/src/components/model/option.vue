<template>
  <div>
    <a-row>
      <a-col :span="3">
        <a-button style="width: 100%;" type="primary" icon="plus" @click="addgroup">新建分组</a-button>
        <a-modal
          title="添加分组"
          v-model="visiblemodel"
          @ok="handleOk"
        >
          <a-input placeholder="输入组名称" v-model="groupname" />
        </a-modal>
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
            <sub-menu v-else :menu-info="item" :key="item.key" :add="add" :send="edit"
:show="showkeypath"/>
          </template>
        </a-menu>
      </a-col>
      <a-col :span="16" style="padding: 0 10px;">
        基本信息
      </a-col>
      <a-col :span="5" style="padding: 0 10px;">
        添加属性
      </a-col>
    </a-row>
    <a-drawer
      :title="menuConfig.group"
      :width="720"
      placement="right"
      :closable="false"
      @close="onClose"
      :visible="visible"
      :wrapStyle="{height: 'calc(100% - 108px)',overflow: 'auto',paddingBottom: '108px'}"
    >
      <a-form :form="form" layout="vertical" hideRequiredMark>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="类型组名">
              <a-input
                v-model="menuConfig.group"
                :disabled="true"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="当前Level">
              <a-input-number
                style="width: 100%;"
                :min="1"
                :max="10"
                v-model="menuConfig.level"
                @change="onChange" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="类型名称">
              <a-input
                v-model="menuConfig.name"
                placeholder="Basic usage"/>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="类型编号">
              <a-input
                v-model="menuConfig.node"
                placeholder="Basic usage"/>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="类型图标">
              <a-select
                v-model="menuConfig.icon"
                placeholder="Please a-s an owner"
              >
                <a-select-opt-group
                  v-for="item in icons"
                  :label="item.title"
                  :key="item.title">
                  <a-select-option
                    v-for="ico in item.icons"
                    :key="ico"
                    :value="ico">
                    <div style="height: 36px"><a-icon :type="ico" /></div>
                  </a-select-option>
                </a-select-opt-group>
              </a-select></a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="应用模板">
              <a-select
                mode="multiple"
                style="width: 100%"
                v-model="menuConfig.template"
                @change="handleSelectChange"
                placeholder="Please select">
                <a-select-option v-for="i in 25" :key="(i + 9).toString(36) + i">{{ (i + 9).toString(36) + i }}</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-button type="primary">编辑属性</a-button>
          </a-col>
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
        <a-button @click="deleteit" type="danger">删除</a-button>
        <a-button @click="deletegroup" type="danger">删除组</a-button>
      </div>
    </a-drawer>
  </div>
</template>

<script>
import SubMenu from './SubMenu1'
import Utils from '@/utils/warehouse'
import icons from './icons'
const Base64 = require('js-base64').Base64
export default {
  components: {
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
    handleSelectChange (i) {
      console.log(i)
    },
    onChange (i) {
      console.log(i)
    },
    submit () {
      if (this.menuConfig.name.length > 0) {
        // var value = {
        //   created: (new Date()).toDateString() + ' ' + (new Date()).toTimeString(),
        //   group: false,
        //   slave: true,
        //   name: this.subgroupame,
        //   parent: this.menuConfig.group
        // }

        this.menuConfig.created = (new Date()).toDateString() + ' ' + (new Date()).toTimeString()
        this.menuConfig.isgroup = false
        this.menuConfig.isslave = true
        var tmp = {
          'key': '/ares/cmdb/model/' + this.menuConfig.group + '/' + this.menuConfig.name,
          'ttl': 1,
          'value': JSON.stringify(this.menuConfig)
        }
        console.log(tmp)
        Utils.setInfo(tmp).then(resp => {
          console.log(resp)
          this.refresh()
          this.$message.success(resp.data.status)
        })
        this.visiblemodel = false
      } else {
        this.$message.error('内容为空，请输入组名')
      }
    },
    deleteit () {
      if (this.menuConfig.name.length > 0) {
        var key = '/ares/cmdb/model/' + this.menuConfig.group + '/' + this.menuConfig.name
        console.log(key)
        Utils.deleteInfo(key).then(resp => {
          console.log(resp)
          this.refresh()
          this.$message.success(resp.data.status)
        })
        this.visiblemodel = false
      } else {
        this.$message.error('内容为空，请输入组名')
      }
    },
    deletegroup () {
      if (this.menuConfig.name.length > 0) {
        var tmp = {
          key: '/ares/cmdb/model/' + this.menuConfig.group,
          recursive: true
        }
        console.log(tmp)
        Utils.delete(tmp).then(resp => {
          console.log(resp)
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
        console.log(data)
        //   var tmp_list = new Map()
        // var decodedData = window.atob(encodedData);
        if (data.data.data.kvs !== undefined) {
          this.list = []
          var tmpData = new Map()
          var num = 0
          data.data.data.kvs.forEach((a) => {
            var tmp = JSON.parse(Base64.decode(a.value))
            this.currentdata.set(tmp.name, tmp)
            if (tmp.isgroup === true && tmp.isslave === false) {
              num += 1
              tmpData.set(tmp.name, { key: tmp.name, title: tmp.name, level: num, children: [] })
            }
          })
          data.data.data.kvs.forEach((b, index) => {
            var tmp1 = JSON.parse(Base64.decode(b.value))
            if (tmp1.isgroup === false && tmp1.isslave === true) {
              if (tmpData.has(tmp1.group)) {
                var t1 = tmpData.get(tmp1.group)
                if (t1.children !== undefined && t1.children.length > 0) {
                  t1.children.push({ key: tmp1.name, title: tmp1.name, icon: tmp1.icon, group: tmp1.group, level: t1.level })
                  tmpData.set(tmp1.group, t1)
                } else {
                  t1.children = [{ key: tmp1.name, title: tmp1.name, icon: tmp1.icon, group: tmp1.group, level: t1.level }]
                  tmpData.set(tmp1.group, t1)
                }
              }
            }
          })
          tmpData.forEach((value, key, map) => {
            this.list.push(value)
          })
        }
        console.log('finished', this.list)
      })
    },
    addgroup () {
      this.visiblemodel = true
    },
    handleOk () {
      if (this.groupname.length > 0) {
        this.menuConfig.created = (new Date()).toDateString() + ' ' + (new Date()).toTimeString()
        this.menuConfig.isgroup = true
        this.menuConfig.isslave = false
        this.menuConfig.group = ''
        this.menuConfig.name = this.groupname
        var tmp = {
          'key': '/ares/cmdb/model/' + this.menuConfig.name,
          'ttl': 1,
          'value': JSON.stringify(this.menuConfig)
        }
        Utils.setInfo(tmp).then(resp => {
          console.log(resp)
          this.refresh()
          this.$message.success(resp.data.status)
        })
        this.visiblemodel = false
      } else {
        this.$message.error('内容为空，请输入组名')
      }
    },
    // handleClick (e) {
    //   console.log('click ', e)
    //   this.current = e.key
    // },
    showDrawer () {
      this.visible = true
    },
    onClose () {
      this.visible = false
      this.keypath = []
    },
    changeTheme (checked) {
      this.theme = checked ? 'dark' : 'light'
    },
    callback (key) {
      console.log(key)
    },
    onSearch (value) {
      console.log(value)
    },
    handleClick (e) {
      console.log('handleClick', e, this.currentdata)
      //   this.menuConfig.name = e.keyPath[0]
      //   this.menuConfig.group = e.keyPath[1]
      this.menuConfig = this.currentdata.get(e.keyPath[0])
      this.keypath = e.keyPath
    //   this.$refs.table.refresh(true)
    },
    edit (name) {
      console.log('edit', name)
      console.log('menuConfig', this.menuConfig)
      console.log('keypath', this.keypath)
      this.menuConfig = this.currentdata.get(name.title)
      this.menuConfig.group = name.group
      this.menuConfig.level = name.level
      this.showDrawer()
    //   this.$refs.table.refresh(true)
    },
    add (name) {
      console.log('add', name)
      console.log('add keypath', this.keypath)
      this.menuConfig = this.currentdata.get(name.title)
      this.menuConfig.group = name.title
      this.menuConfig.name = ''
      this.menuConfig.node = ''
      this.menuConfig.icon = ''
      this.menuConfig.template = []
      this.menuConfig.level = name.level
      console.log('add menuConfig', this.menuConfig)
      this.showDrawer()
    //   this.$refs.table.refresh(true)
    },
    handleClick2 (e) {
      console.log('handleClick222', e, this.keyPath)
      //   console.log(e.target.outerText)
      // menu 默认为2级目录 超过数据会出错
      // menu 默认也不能为1级信息
      this.menuConfig = this.currentdata.get(e.target.outerText)
      if (this.keypath.length === 2) {
        // 点击非叶子节点
        console.log('1', e.target.outerText, this.keypath)
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
