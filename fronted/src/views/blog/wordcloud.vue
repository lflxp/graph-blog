<template>
  <div id="app">
    <!-- <wordcloud
      :data="defaultWords"
      nameKey="name"
      valueKey="value"
      :color="myColors"
      :showTooltip="true"
      :wordClick="wordClickHandler">
    </wordcloud>-->
    <a-row>
      <a-col :span="8" :offset="8">
        <a-auto-complete
          class="global-search"
          size="large"
          style="width: 100%"
          @select="onSelect"
          @search="handleSearch"
          placeholder="输入文章属性"
          optionLabelProp="text"
        >
          <template slot="dataSource">
            <a-select-option v-for="item in dataSource" :key="item.category" :text="item.category">
              {{item.query}} 在
              <a
                :href="`https://s.taobao.com/search?q=${item.query}`"
                target="_blank"
                rel="noopener noreferrer"
              >{{item.category}}</a>
              区块中
              <span class="global-search-item-count">约 {{item.count}} 个结果</span>
            </a-select-option>
          </template>
          <a-input>
            <a-button slot="suffix" class="search-btn" size="large" type="primary">
              <a-icon type="search"/>搜索
            </a-button>
          </a-input>
        </a-auto-complete>
      </a-col>
    </a-row>
    <a-row>
      <a-col>
        <cloud 
          class="cloud"
          :onWordClick="handClick"
          :data="words" 
          :padding="padding"
          :font="font"
          :rotate="rotate"
          :fontSizeMapper="fontSizeMapper"/>
      </a-col>
    </a-row>
  </div>
</template>

<script>
import wordcloud from 'vue-wordcloud'
import Cloud from 'vue-d3-cloud'
import Utils from '@/utils/service'

export default {
  name: 'app',
  components: {
    wordcloud,
    Cloud
  },
  methods: {
    handClick (e) {
      alert(e)
      console.log('handClick', e)
    },
    wordClickHandler(name, value, vm) {
      console.log('wordClickHandler', name, value, vm)
    },
    onSelect(value) {
      console.log('onSelect', value)
    },

    handleSearch(value) {
      this.dataSource = value ? this.searchResult(value) : []
    },

    getRandomInt(max, min = 0) {
      return Math.floor(Math.random() * (max - min + 1)) + min
    },

    searchResult(query) {
      return new Array(this.getRandomInt(5))
        .join('.')
        .split('.')
        .map((item, idx) => ({
          query,
          category: `${query}${idx}`,
          count: this.getRandomInt(200, 100)
        }))
    }
  },
  data() {
    return {
      min: 0,
      max: 65,
      padding: 0,
      font: 'Helvetica',
      dataSource: [],
      myColors: ['#1f77b4', '#629fc9', '#yellow', '#red'],
      defaultWords: [
        {
          name: 'Cat',
          value: 26
        },
        {
          name: 'fish',
          value: 19
        },
        {
          name: 'things',
          value: 18
        },
        {
          name: 'look',
          value: 16
        },
        {
          name: 'two',
          value: 15
        },
        {
          name: 'fun',
          value: 9
        },
        {
          name: 'know',
          value: 9
        },
        {
          name: 'good',
          value: 9
        },
        {
          name: 'play',
          value: 6
        }
      ],
      words: [
        { text: 'Vue', value: 1000 },
        { text: 'js', value: 200 },
        { text: 'is', value: 800 },
        { text: 'very cool', value: 1000 },
        { text: 'lunch', value: 100 }
      ],
      fontSizeMapper: word => Math.log2(word.value) * 5
    }
  },
  computed: {
		rotate: function() {
			let newMin = this.min
			let newMax = this.max
			return word => Math.random() * (newMax - newMin) + newMin
    },
    wordCount: function() {
			if(!this.words)
				return []
			let occurences = this.words.split(' ').reduce((allNames, name) => { 
				if (name in allNames) {
					allNames[name]++;
				} else {
					allNames[name] = 1;
				}
				return allNames;
			}, {});
			
			let occurencesCount = []
			
			for(var text in occurences) {
				let obj = {}
				obj.text = text;
				obj.value = occurences[text]
				occurencesCount.push(obj)
			}
			return occurencesCount
		}
  },
  mounted() {
    var cql = { cql: 'MATCH (n) WHERE rand() <= 0.5 return n' }
    console.log(cql)
    Utils.getCql(cql).then(data => {
      console.log('dddd', data)
      if (data.data.nodes !== undefined && data.data.nodes.length > 0) {
        var tmp = new Map()
        data.data.nodes.forEach(info => {
          // if (info.props.title !== undefined) {
          //   if (tmp.has(info.props.title)) {
          //     var t1 = tmp.get(info.props.title)
          //     t1.value += 1
          //     tmp.set(info.props.title, t1)
          //   } else {
          //     var t2 = {
          //       name: info.props.title,
          //       value: 1
          //     }
          //     tmp.set(info.props.title, t2)
          //   }
          // }
          // if (info.labels[0] !== undefined) {
          //   if (tmp.has(info.labels[0])) {
          //     var t1 = tmp.get(info.labels[0])
          //     t1.value += 1
          //     tmp.set(info.labels[0], t1)
          //   } else {
          //     var t2 = {
          //       name: info.labels[0],
          //       value: 1
          //     }
          //     tmp.set(info.labels[0], t2)
          //   }
          // }
          console.log('info', info)
          if (info.props.title !== undefined) {
            if (tmp.has(info.props.title)) {
              var t1 = tmp.get(info.props.title)
              t1.value += 1
              tmp.set(info.props.title, t1)
            } else {
              var t2 = {
                text: info.props.title,
                value: 100
              }
              tmp.set(info.props.title, t2)
            }
          }
        })
        tmp.forEach((value, key, map) => {
          console.log(value, key, map)
          // this.defaultWords.push(value)
          this.words.push(value)
        })
        console.log(this.words)
      }
    })
  }
}
</script>

<style>
.global-search-wrapper {
  padding-right: 50px;
}

.global-search {
  width: 100%;
}

.global-search.ant-select-auto-complete .ant-select-selection--single {
  margin-right: -46px;
}

.global-search.ant-select-auto-complete .ant-input-affix-wrapper .ant-input:not(:last-child) {
  padding-right: 62px;
}

.global-search.ant-select-auto-complete .ant-input-affix-wrapper .ant-input-suffix {
  right: 0;
}

.global-search.ant-select-auto-complete .ant-input-affix-wrapper .ant-input-suffix button {
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

.global-search-item-count {
  position: absolute;
  right: 16px;
}

.cloud svg {
  height: 800px;
  /* width: 1200px; */
  display: block;
  margin-left: auto;
  margin-right: auto;
}

/* .cloud {
  height: 800px;
  width: 1200px;
} */
</style>
