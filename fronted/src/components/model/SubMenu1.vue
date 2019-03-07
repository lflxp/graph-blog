<template>
  <a-sub-menu
    :key="menuInfo.key"
    v-bind="$props"
    v-on="$listeners">
    <span slot="title">
      <a-popover placement="left" v-if="showadd">
        <template slot="content">
          <a-button type="dashed" @click="add(menuInfo)">新增</a-button>
        </template>
        <a-icon type="mail" />
        <span>{{ menuInfo.title }}</span>
      </a-popover>
      <div v-else>
        <a-icon type="mail" />
        <span>{{ menuInfo.title }}</span>
      </div>
    </span>
    <template v-for="item in menuInfo.children">
      <a-menu-item
        v-if="!item.children"
        :key="item.key">
        <a-popover placement="right">
          <template slot="content">
            <a-button type="dashed" @click="send(item)">编辑</a-button>
          </template>
          <a-icon :type="item.icon" />
          <span>{{ item.title }}</span>
        </a-popover>
      </a-menu-item>
      <sub-menu
        v-else
        :key="item.key"
        :menu-info="item"/>
    </template>
  </a-sub-menu>
</template>
<script>
import { Menu } from 'ant-design-vue'
export default {
  name: 'SubMenu',
  // must add isSubMenu: true
  isSubMenu: true,
  props: {
    ...Menu.SubMenu.props,
    // Cannot overlap with properties within Menu.SubMenu.props
    menuInfo: {
      type: Object,
      default: () => ({})
    },
    send: {
      type: Function
    },
    show: {
      type: Boolean
    },
    add: {
      type: Function
    },
    showadd: {
      type: Boolean,
      default: true
    }
  }
}
</script>
