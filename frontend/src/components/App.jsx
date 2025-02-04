import { useState } from 'react';
import './App.css';
import Header from './Header';
import Todo from './Todo';

function App() {
  const [theme, setTheme] = useState(false);
  const changeTheme = (t) => {
    setTheme(t);
  }
  return (
    <div className="App" style={{background: `${theme === true ? "white" : "var(--body-color)"}`}}>
      <Header theme={theme}/>
      <Todo changeTheme={changeTheme}/>;
    </div>
  );
}

export default App;
