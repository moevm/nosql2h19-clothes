import React from 'react';
// import ReactDOM from 'react-dom';
import './index.css';
import Form from './Form';
import Table from './Table';
import AddSmth from './AddSmth.js';
import Home from './Home';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';

function App() {
  return (
    <Router>  
      <div className="App">
        <Switch>
          <Route path="/" exact component={Form} />
          <Route path="/table" component={Table} />
          <Route path="/add_category" component={AddSmth} />
          <Route path="/add_place" component={AddSmth} />
          <Route path="/home" exact component={Home} />
          <Route path="/home/day" component={Table} />
        </Switch>
      </div>
    </Router>
  );
}
// import { format } from 'path';

export default App;
