new Vue({
  el: '#github',
  data:{
    repos:[],
    github_id:location.search.substr(1)
  },
  mounted () {
    axios
      .get(`https://api.github.com/users/${this.github_id}/repos`)
      .then(response => (this.repos = response.data))
  }
})

new Vue({
  el: '#qiita',
  data:{
    articles:[],
    qiita_id:""
  },
  methods: {
    qiita_button:function() {
      axios
      .get(`https://qiita.com/api/v2/users/${this.qiita_id}/items`)
      .then(response => (this.articles = response.data))
    }
  }
})

new Vue({
  el: "#kusa",
  data:{
    kusa:`http://localhost:8080/user/${location.search.substr(1)}/kusa`
  }
})
