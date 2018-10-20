class TodoApp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {items: [], Github: '', Qiita: ''};
  }
  
  render() {
    return (
        <div className='container'>
        <ul></ul>
          <div className='nav'>管理画面</div>
          <img className='icon' src='../image.png'></img>
          
          <form>
            <div className="form-group">
              <label for="exampleInputEmail1">GitHub Email Address</label>
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
              <label for="exampleInputEmail1">Qiita Email Adress</label>
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
            <button type='button' className="btn btn-dark" onClick={this.handleSubmit.bind(this)}>連携</button>
          </form>
          <TodoList items={this.state.items} />
      </div>
    );
  }

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
  <TodoApp />,
  document.getElementById('content')
);
