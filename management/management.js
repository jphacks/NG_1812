class TodoApp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {items: [], Github: '', Qiita: ''};
  }
  render() {
    return (
        <div className='container'>
          <div className='nav'>管理画面</div>
          <form>
            <ul>
              <li>GitHub ID</li>
                <li>
                  <input onChange={this.handleChangeGithub.bind(this)} value={this.state.Github} />
                </li>
                <li>Qiita ID</li>
                <li>
                  <input onChange={this.handleChangeQiita.bind(this)} value={this.state.Qiita} />
                </li>
              </ul>
              <button type='button' onClick={this.handleSubmit.bind(this)}>連携</button>
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
    fetch('https://httpbin.org/status/tanaka', {
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
