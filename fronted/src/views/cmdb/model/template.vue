<template>
  <div>
    <a-tabs defaultActiveKey="1" @change="callback">
      <a-tab-pane tab="模型属性" key="1">
        <templated></templated>
      </a-tab-pane>
      <a-tab-pane tab="模型关系" key="2" forceRender>
        <template-relation v-if="show"></template-relation>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script>
import Templated from '@/components/model/template'
import TemplateRelation from '@/components/model/templateRelation'
export default {
  components: {
    Templated,
    TemplateRelation
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
