new Vue({
    el: '#github',
    data () {
      return {
        repos:[]
      }
    },
    mounted () {
      axios
        .get('https://api.github.com/users/tmk815/repos')
        .then(response => (this.repos = response.data))
    }
})

new Vue({
    el: '#qiita',
    data () {
      return {
        articles:[]
      }
    },
    mounted () {
      axios
        .get('https://qiita.com/api/v2/users/tmk815/items')
        .then(response => (this.articles = response.data))
    }
  })