class Management extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
      Github: '',
      Qiita: '',
      git_user_url: 'https://api.github.com/users/',
      icon_url: './src/default.png'
    };
  }

  componentDidMount = () => {
    // get query list from url
    const getQuery = () => {
      if(window.location.search === "") return;
      const variables = window.location.search.split("?")[1].split("&");
      const obj = {};
      variables.forEach(function(v, i) {
          const variable = v.split("=");
          obj[variable[0]] = Number(variable[1]);
      });
      return obj;
    }

    const code = getQuery().code;
    if(code != ""){
      document.cookie = `accessToken=${code}`;
    }
    //this.getUserIcon(this.state.git_user_url);
  };

  getUserIcon = (url) => {
    fetch(url, {
      method: 'GET'
    }).then(response => {
      return response.json();
    })
    .then(text => {
      let icon = text.avatar_url;
      this.setState({icon_url: icon});
    });
  }

  handleChangeGithub(e) {
    this.setState({Github: e.target.value});
  }


  handleChangeQiita(e) {
    this.setState({Qiita: e.target.value});
  }


  handleSubmit = () => {
    const ID = {
      GithubID: this.state.Github,
      QiitaID: this.state.Qiita
    };
    this.getUserIcon(this.state.git_user_url + ID.GithubID);
    fetch('https://httpbin.org/status/dummy', {
      method: 'POST',
      body: ID,
      })
      .then(response => {
      if (!response.ok) {
        window.alert('失敗');
      }
    });
  }
  
  render() {
    return (
      <div className='container'>
      <ul></ul>
        <div className='nav'>管理画面</div>
        <div id='profile'>
          <img className='icon' src={this.state.icon_url}></img>
          <form className='bio' >
            <div className="form-group">
              <label htmlFor="bio">プロフィール欄</label>
              <textarea className="form-control" rows="3"></textarea>
            </div>
          </form>
          </div>
        <form>
          <div className="col-auto">
            <label htmlFor="inlineFormInputGroup">GitHub Username</label>
            <div className="input-group mb-2">
            <input
              type="text"
              className="form-control"
              placeholder="Enter GitHub username"
              onChange={this.handleChangeGithub.bind(this)}
              value={this.state.Github}
            />
            </div>
          </div>
          <div className="col-auto">
            <label htmlFor="inlineFormInputGroup">Qiita Username</label>
            <div className="input-group mb-2">
            <input
              type="text"
              className="form-control"
              placeholder="Enter Qiita username"
              onChange={this.handleChangeQiita.bind(this)}
              value={this.state.Qiita}
            />
            </div>
          </div>
          <button type='button' className="btn btn-dark" onClick={this.handleSubmit.bind(this)}>連携</button>
        </form>
        <TodoList items={this.state.items} />
      </div>
    );
  }
}


const TodoList = (props) => {
  return (
    <ul>
      {props.items.map((item, index) => (
        <li key={index}>{item}</li>
      ))}
    </ul>
  );
}

ReactDOM.render(
  <Management />,
  document.getElementById('content')
);