class Management extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
      Github: '',
      Qiita: '',
      git_user_url: 'https://api.github.com/users/tmk815',
      icon_url: ''
    };
  }

  componentDidMount = () => {
    fetch(this.state.git_user_url, {
      method: 'GET'
    }).then(response => {
      return response.json();
    })
    .then(text => {
      let icon = text.avatar_url;
      this.setState({icon_url: icon});
    });
  };

  handleChangeGithub(e) {
    this.setState({Github: e.target.value});
  }


  handleChangeQiita(e) {
    this.setState({Qiita: e.target.value});
  }


  handleSubmit(e) {
    const ID = {
      GithubID: this.state.Github,
      QiitaID: this.state.Qiita
    };
    fetch('https://httpbin.org/status/dummy', {
      method: 'POST',
      body: ID,
      })
      .then(response => {
      if (!response.ok) {
        window.alert(ID.GithubID);
      }
    });
  }
  
  render() {
    return (
      <div className='container'>
      <ul></ul>
        <div className='nav'>管理画面</div>
        <img className='icon' src={this.state.icon_url}></img>
        <form>
          {/*
          <div className="form-group">
            <label htmlFor="exampleInputEmail1">GitHub Email Address</label>
            <input
              type="email"
              className="form-control"
              id="exampleInputEmail1"
              aria-describedby="emailHelp"
              placeholder="Enter GitHub email"
              onChange={this.handleChangeGithub.bind(this)}
              value={this.state.Github}
            />
          </div>
          <div className="form-group">
            <label htmlFor="exampleInputEmail1">Qiita Email Adress</label>
            <input
              type="email"
              className="form-control"
              id="exampleInputEmail1"
              aria-describedby="emailHelp"
              placeholder="Enter Qiita email"
              onChange={this.handleChangeQiita.bind(this)}
              value={this.state.Qiita}
            />
          </div>
          */}
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
