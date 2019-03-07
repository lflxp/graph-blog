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
          <a-input placeholder="输入组名称" v-model="groupame" />
        </a-modal>
        <a-menu
          style="width: 100%;"
          :defaultSelectedKeys="['1']"
          :defaultOpenKeys="['2']"
          :selectedKeys="selects"
          @click="handleClick"
          @dblclick="handleClick2"
          mode="inline"
          theme="dark"
          :inlineCollapsed="collapsed"
        >
          <template v-for="item in list">
            <a-menu-item v-if="!item.children" :key="item.key">
              <a-icon type="pie-chart" />
              <span>{{ item.title }}</span>
            </a-menu-item>
            <sub-menu v-else :menu-info="item" :key="item.key"/>
          </template>
        </a-menu>
      </a-col>
      <a-col :span="21" style="padding: 0 10px;">
        <a-row type="flex" justify="start">
          <a-col :span="6">
            <a-input-search
              placeholder="搜索"
              size="default"
              style="width: 200px"
              @search="onSearch"
            />
          </a-col>
          <a-col :span="18">
            <div style="float:right;">
              <a-icon type="up-square" />
              <a-divider type="vertical" />
              <a-icon type="down-square" />
              <a-divider type="vertical" />
              <a-icon type="file-text" />Text
              <a-divider type="vertical" />
              <a-icon type="link" /><a href="#">Link</a>
              <a-divider type="vertical" />
              <a-icon type="link" /><a href="#">Link</a>
              <a-divider type="vertical" />
              <a-dropdown>
                <a class="ant-dropdown-link" href="#">
                  更多 <a-icon type="down" />
                </a>
                <a-menu slot="overlay">
                  <a-menu-item>
                    <a href="javascript:;">1st menu item</a>
                  </a-menu-item>
                  <a-menu-item>
                    <a href="javascript:;">2nd menu item</a>
                  </a-menu-item>
                  <a-menu-item>
                    <a href="javascript:;">3rd menu item</a>
                  </a-menu-item>
                </a-menu>
              </a-dropdown>
              <a-button type="primary">添加配置</a-button>
              <a-button type="primary">认领</a-button>
            </div>
          </a-col>
        </a-row>
        <a-row>
          <a-table ref="table" :columns="columns" :dataSource="data" style="padding: 10px 0 0 0;">
            <a slot="name" slot-scope="text" href="javascript:;">{{ text }}</a>
            <span slot="customTitle"><a-icon type="smile-o" /> Name</span>
            <span slot="tags" slot-scope="tags">
              <a-tag v-for="tag in tags" color="blue" :key="tag">{{ tag }}</a-tag>
            </span>
            <span slot="action" slot-scope="text, record">
              <a href="javascript:;">Invite 一 {{ record.name }}</a>
              <a-divider type="vertical" />
              <a href="javascript:;">Delete</a>
              <a-divider type="vertical" />
              <a href="javascript:;" class="ant-dropdown-link">
                More actions <a-icon type="down" />
              </a>
            </span>
          </a-table>
        </a-row>
      </a-col>
    </a-row>
    <a-drawer
      :title="addtitle"
      placement="right"
      :closable="false"
      @close="onClose"
      :visible="visible"
    >
      <a-row>
        <a-col>
          <a-input placeholder="输入详细配置名" size="small" v-model="subgroupame" />
        </a-col>
      </a-row>
      <a-row style="padding: 10px 0;float:right;">
        <a-col>
          <a-button type="dashed" @click="cancel">取消</a-button><a-button type="primary" @click="addsub">新增</a-button>
        </a-col>
      </a-row>
    </a-drawer>
  </div>
</template>
<script>
/**
 * configPath: /ares/cmdb/warehouse/group/name
 */
import SubMenu from './SubMenu'
import Utils from '@/utils/warehouse'
const Base64 = require('js-base64').Base64
const columns = [{
  dataIndex: 'name',
  key: 'name',
  slots: { title: 'customTitle' },
  scopedSlots: { customRender: 'name' }
}, {
  title: 'Age',
  dataIndex: 'age',
  key: 'age'
}, {
  title: 'Address',
  dataIndex: 'address',
  key: 'address'
}, {
  title: 'Tags',
  key: 'tags',
  dataIndex: 'tags',
  scopedSlots: { customRender: 'tags' }
}, {
  title: 'Action',
  key: 'action',
  scopedSlots: { customRender: 'action' }
}]

const data = [{
  key: '1',
  name: 'John Brown',
  age: 32,
  address: 'New York No. 1 Lake Park',
  tags: ['nice', 'developer']
}, {
  key: '2',
  name: 'Jim Green',
  age: 42,
  address: 'London No. 1 Lake Park',
  tags: ['loser']
}, {
  key: '3',
  name: 'Joe Black',
  age: 32,
  address: 'Sidney No. 1 Lake Park',
  tags: ['cool', 'teacher']
}]
export default {
  components: {
    'sub-menu': SubMenu
  },
  data () {
    return {
      subgroupame: '',
      groupame: '',
      visiblemodel: false,
      visible: false,
      addtitle: '',
      keypath: [],
      current: '1',
      theme: 'dark',
      columns,
      data,
      selects: [],
      collapsed: false,
      list: []
    }
  },
  mounted () {
    this.refresh()
  },
  methods: {
    addsub () {
      if (this.subgroupame.length > 0) {
        var value = {
          created: (new Date()).toDateString() + ' ' + (new Date()).toTimeString(),
          group: false,
          slave: true,
          name: this.subgroupame,
          parent: this.addtitle
        }
        var tmp = {
          'key': '/ares/cmdb/warehouse/' + this.addtitle + '/' + this.subgroupame,
          'ttl': 1,
          'value': JSON.stringify(value)
        }
        console.log(tmp)
        Utils.setInfo(tmp).then(resp => {
          console.log(resp)
          //   this.refresh()
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
        'key': '/ares/cmdb/warehouse',
        'recursive': true
      }
      Utils.getInfo(se).then(data => {
        console.log(data)
        //   var tmp_list = new Map()
        // var decodedData = window.atob(encodedData);
        if (data.data.data.kvs !== undefined) {
          this.list = []
          var tmpData = new Map()
          data.data.data.kvs.forEach((a) => {
            var tmp = JSON.parse(Base64.decode(a.value))
            console.log(Base64.decode(a.key), Base64.decode(a.value), tmp)
            if (tmp.group === true && tmp.slave === false) {
              tmpData.set(tmp.name, { key: tmp.name, title: tmp.name, children: [] })
            }
          })
          data.data.data.kvs.forEach((b) => {
            var tmp1 = JSON.parse(Base64.decode(b.value))
            if (tmp1.group === false && tmp1.slave === true) {
              if (tmpData.has(tmp1.parent)) {
                console.log('parent', tmpData, tmp1)
                var t1 = tmpData.get(tmp1.parent)
                console.log('t1', t1)
                if (t1.children !== undefined && t1.children.length > 0) {
                  t1.children.push({ key: tmp1.name, title: tmp1.name })
                  tmpData.set(tmp1.parent, t1)
                } else {
                  t1.children = [{ key: tmp1.name, title: tmp1.name }]
                  tmpData.set(tmp1.parent, t1)
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
      if (this.groupame.length > 0) {
        var value = {
          created: (new Date()).toDateString() + ' ' + (new Date()).toTimeString(),
          group: true,
          slave: false,
          name: this.groupame,
          parent: ''
        }
        var tmp = {
          'key': '/ares/cmdb/warehouse/' + this.groupame,
          'ttl': 1,
          'value': JSON.stringify(value)
        }
        console.log(tmp)
        Utils.setInfo(tmp).then(resp => {
          console.log(resp)
          //   this.refresh()
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
      console.log('handleClick', e)
      this.keypath = e.keyPath
    //   this.$refs.table.refresh(true)
    },
    handleClick2 (e) {
      console.log('handleClick222', e, this.keyPath)
      console.log(e.target.outerText)
      // menu 默认为2级目录 超过数据会出错
      // menu 默认也不能为1级信息
      if (this.keypath.length > 1) {
        // 点击非叶子节点
        if (e.target.outerText === this.keypath[1]) {
          console.log('1', e.target.outerText, this.keypath)
          this.addtitle = e.target.outerText
        } else {
          console.log('2', e.target.outerText, this.keypath)
          this.addtitle = this.keypath[1]
        }
      } else if (this.keypath.length === 1) {
        if (e.target.outerText === this.keypath[0]) {
          console.log('1.1', e.target.outerText, this.keypath)
          this.addtitle = e.target.outerText
        } else {
          console.log('2.1', e.target.outerText, this.keypath)
          this.addtitle = this.keypath[0]
        }
      } else {
        console.log('3', e.target.outerText, this.keypath)
        this.addtitle = e.target.outerText
      }
      this.showDrawer()
    //   this.$refs.table.refresh(true)
    }
  }
}

</script>
