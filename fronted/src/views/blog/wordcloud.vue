<template>
  <div id="app">
      <!-- <wordcloud
      :data="defaultWords"
      nameKey="name"
      valueKey="value"
      :color="myColors"
      :showTooltip="true"
      :wordClick="wordClickHandler">
      </wordcloud> -->
      <cloud :data="words" :fontSizeMapper="fontSizeMapper" />
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
    wordClickHandler(name, value, vm) {
      console.log('wordClickHandler', name, value, vm);
    }
  },
  data() {
    return {
      myColors: ['#1f77b4', '#629fc9', '#yellow', '#red'],
      defaultWords: [{
          "name": "Cat",
          "value": 26
        },
        {
          "name": "fish",
          "value": 19
        },
        {
          "name": "things",
          "value": 18
        },
        {
          "name": "look",
          "value": 16
        },
        {
          "name": "two",
          "value": 15
        },
        {
          "name": "fun",
          "value": 9
        },
        {
          "name": "know",
          "value": 9
        },
        {
          "name": "good",
          "value": 9
        },
        {
          "name": "play",
          "value": 6
        }
      ],
      words: [
          { text: 'Vue', value: 1000 },
          { text: 'js', value: 200 },
          { text: 'is', value: 800 },
          { text: 'very cool', value: 1000 },
          { text: 'lunch', value: 100 },
      ],
      fontSizeMapper: word => Math.log2(word.value) * 5
    }
  },
  mounted() {
    var cql = {cql: 'MATCH (n) WHERE rand() <= 0.5 return n'}
    Utils.getCql(cql).then(data=>{
      console.log(data)
      if (data.data.nodes !== undefined && data.data.nodes.length > 0) {
        var tmp = new Map()
        data.data.nodes.forEach(info=>{
          // if (info.props.title !== undefined) {
          //   if (tmp.has(info.props.title)) {
          //     var t1 = tmp.get(info.props.title)
          //     t1.value += 1
          //     tmp.set(info.props.title,t1)
          //   } else {
          //     var t2 = {
          //       name: info.props.title,
          //       value: 1
          //     }
          //     tmp.set(info.props.title,t2)
          //   }
          // }
          // if (info.labels[0] !== undefined) {
          //   if (tmp.has(info.labels[0])) {
          //     var t1 = tmp.get(info.labels[0])
          //     t1.value += 1
          //     tmp.set(info.labels[0],t1)
          //   } else {
          //     var t2 = {
          //       name: info.labels[0],
          //       value: 1
          //     }
          //     tmp.set(info.labels[0],t2)
          //   }
          // }
          if (info.props.title !== undefined) {
            if (tmp.has(info.props.title)) {
              var t1 = tmp.get(info.props.title)
              t1.value += 1
              tmp.set(info.props.title,t1)
            } else {
              var t2 = {
                text: info.props.title,
                value: 100
              }
              tmp.set(info.props.title,t2)
            }
          }
        })
        tmp.forEach((value,key,map)=>{
          console.log(value,key,map)
          // this.defaultWords.push(value)
          this.words.push(value)
        })
        console.log(this.words)
      }
    })
  },
}
</script>