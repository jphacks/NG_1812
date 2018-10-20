new Vue({
  el: '#github',
  data:{
    repos:[],
    github_id:""
  },
  methods: {
    github_button:function() {
      axios
      .get(`https://api.github.com/users/${this.github_id}/repos`)
      .then(response => (this.repos = response.data))
    }
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