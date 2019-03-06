<template>
  <div :style="!$route.meta.pageHeader ? 'margin: -24px -24px 0px;' : null">
    <!-- pageHeader , route meta hideHeader:true on hide -->
    
    <div class="content">
      <div :class="['page-header-index-wide']">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
import PageHeader from './PageHeader'

export default {
  name: 'LayoutContent',
  components: {
    PageHeader
  },
  // ['desc', 'logo', 'title', 'avatar', 'linkList', 'extraImage']
  props: {
    desc: {
      type: String,
      default: null
    },
    logo: {
      type: String,
      default: null
    },
    title: {
      type: String,
      default: null
    },
    avatar: {
      type: String,
      default: null
    },
    linkList: {
      type: Array,
      default: null
    },
    extraImage: {
      type: String,
      default: null
    },
    search: {
      type: Boolean,
      default: false
    },
    tabs: {
      type: Object,
      default: () => ({})
    }
  },
  methods: {
  }
}
</script>

<style lang="less" scoped>
  .content {
    margin: 24px 24px 0;

    .link {
      margin-top: 16px;

      &:not(:empty) {
        margin-bottom: 16px;
      }
      a {
        margin-right: 32px;
        height: 24px;
        line-height: 24px;
        display: inline-block;

        i {
          font-size: 24px;
          margin-right: 8px;
          vertical-align: middle;
        }
        span {
          height: 24px;
          line-height: 24px;
          display: inline-block;
          vertical-align: middle;
        }
      }
    }
  }
  .page-menu-search {
    text-align: center;
    margin-bottom: 16px;
  }
  .page-menu-tabs {
    margin-top: 48px;
  }
</style>
