<template>
  <div>
    <a-tabs defaultActiveKey="1" @change="callback">
      <a-tab-pane tab="类型属性" key="1">
        <item></item>
      </a-tab-pane>
      <a-tab-pane tab="类型关系" key="2" forceRender>
        <option-relation v-if="show"></option-relation>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script>
import Item from '@/components/model/option'
import optionRelation from '@/components/model/optionRelation'
export default {
  components: {
    Item,
    optionRelation
  },
  data () {
    return {
      show: true
    }
  },
  methods: {
    callback (key) {
      if (key === '2') {
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
      }
    }
  }
}
</script>
