<template functional>
  <a-sub-menu
    :key="props.menuInfo.key"
  >
    <span slot="title">
      <a-icon type="mail" /><span>{{ props.menuInfo.title }}</span>
    </span>
    <template v-for="item in props.menuInfo.children">
      <a-menu-item
        v-if="!item.children"
        :key="item.key"
      >

        <a-popover placement="right">
          <template slot="content">
            <a-button type="dashed" v-on:click="send(item)">编辑</a-button>
          </template>
          <a-icon type="pie-chart" />
          <span>{{ item.title }}</span>
        </a-popover>
      </a-menu-item>
      <sub-menu
        v-else
        :key="item.key"
        :menu-info="item"
      />
    </template>
  </a-sub-menu>
</template>
<script>
export default {
  props: ['menuInfo'],
  methods: {
    send (item) {
      console.log('111111111111', item)
      this.$emit('listenson', item)
    }
  }
}
</script>
