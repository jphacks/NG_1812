new Vue({
  el: '#github',
  data:{
    repos:[],
    github_id:location.search.substr(1)
  },
  mounted () {
    axios
      .get(`http://ec2-18-191-90-196.us-east-2.compute.amazonaws.com:8080/repos/user/${this.github_id}`)
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
    qiita_button:async function() {
      let page = 1;
      let data = []
      do{
        data = (await axios.get(`https://qiita.com/api/v2/users/${this.qiita_id}/items?page=${page}`)).data
        this.articles = this.articles.concat(data)
        page++
      }while(data.length != 0)
      this.articles.crea
    }
  }
})

new Vue({
  el: "#kusa",
  data:{
    kusa:`http://localhost:8080/user/${location.search.substr(1)}/kusa`
  }
})
