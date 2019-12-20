import React from 'react';
// import ReactDOM from 'react-dom';
import './index.css';
import Form from './Form';
import Places from './Places';
import Categories from './Categories';
import AddPlace from './AddPlace.js';
import AddCategories from './AddCategories.js';
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
import Clothes from "./Clothes";
import AddClothe from "./AddClothe";

function App() {
  return (
    <Router>  
      <div className="App">
        <Switch>
          <Route path="/" exact component={Form} />
          <Route path="/home/places" exact component={Places} />
          <Route path="/home/places/edit" component={EditPlace} />
          <Route path="/home/categories" exact component={Categories} />
          <Route path="/home/categories/add" exact component={AddCategories} />
          <Route path="/home/categories/show" exact component={ShowCat} />
          <Route path="/home/categories/show/add" exact component={AddObj} />
          <Route path="/home/places/add" component={AddPlace} />
          <Route path="/home" exact component={Home} />
          <Route path="/home/day" component={Day} />
          <Route path="/home/make" component={Make} />
          <Route path="/admin" exact component={Admin} />
          <Route path="/admin/add" exact component={AddUser} />
          <Route path="/admin/edit" exact component={EditUser} />
          <Route path="/home/clothes" exact component={Clothes} />
          <Route path="/home/clothes/add" exact component={AddClothe} />
        </Switch>
      </div>
    </Router>
  );
}
// import { format } from 'path';

export default App;
