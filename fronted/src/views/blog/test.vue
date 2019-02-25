<template>
  <div>
    <a-button type="primary" @click="showDrawer">
      Open
    </a-button>
    <a-drawer
      title="Basic Drawer"
      placement="right"
      :closable="false"
      @close="onClose"
      :visible="visible"
    >
      <p>Some contents...</p>
      <p>Some contents...</p>
      <p>Some contents...</p>
    </a-drawer>
    <a-divider type="vertical" />
    <a href="https://vue.ant.design/components/drawer-cn/">官网</a>
    <a-divider/>
    <a-popconfirm title="Are you sure delete this task?" @confirm="confirm" @cancel="cancel" okText="Yes" cancelText="No">
        <a href="#">Delete</a>
    </a-popconfirm>
  </div>
</template>
<script>
export default {
  data () {
    return {
      visible: false
    }
  },
  methods: {
    showDrawer () {
      this.visible = true
    },
    onClose () {
      this.visible = false
    },
    confirm (e) {
      console.log(e)
      this.$message.success('Click on Yes')
    },
    cancel (e) {
      console.log(e)
      this.$message.error('Click on No')
    }
  }
}
</script>
