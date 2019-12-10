import React from 'react';
// import ReactDOM from 'react-dom';
import './index.css';
import Form from './Form';
import Table from './Table';
import AddSmth from './AddSmth.js';
import Home from './Home';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';
import Make from './Make';
import Day from './Day';
import Admin from './Admin';
import AddUser from './AddUser';
import EditUser from './EditUser';
import EditPlace from './EditPlace';
import ShowCat from './ShowCat';
import AddObj from './AddObj';

function App() {
  return (
    <Router>  
      <div className="App">
        <Switch>
          <Route path="/" exact component={Form} />
          <Route path="/home/places" exact component={Table} />
          <Route path="/home/places/edit" component={EditPlace} />
          <Route path="/home/categories" exact component={Table} />
          <Route path="/home/categories/show" exact component={ShowCat} />
          <Route path="/home/categories/show/add" exact component={AddObj} />
          <Route path="/add" component={AddSmth} />
          <Route path="/home" exact component={Home} />
          <Route path="/home/day" component={Day} />
          <Route path="/home/make" component={Make} />
          <Route path="/admin" exact component={Admin} />
          <Route path="/admin/add" exact component={AddUser} />
          <Route path="/admin/edit" exact component={EditUser} />
        </Switch>
      </div>
    </Router>
  );
}
// import { format } from 'path';

export default App;
