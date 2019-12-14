import React, {Component} from 'react';
import axios from 'axios';
import {Link} from 'react-router-dom';
import './Table.css';

let endpoint = "http://localhost:5000";

class Categories extends Component {
  constructor(props) {
      super(props);

      this.state = {
          categories: []
      }
  }
  componentDidMount() {
      this.getCategories();
  }
    
  getCategories = () => {
    axios
      .get(endpoint + "/api/home/categories")
      .then(res => {
          console.log(res);
          this.setState({
              categories: res.data.map(item => {
                  return item['name'];
              })
          })
      })
    //   fetch("http://localhost:5000/api/home/places")
    //       .then(data => console.log(data));
  }

  render() {
    let rows = [];
    this.state.categories.forEach((item, i) => {
        rows.push(<tr key={i}><td>{item}</td>
        <td><button className="green">see</button><button className="red">del</button></td></tr>)
    })
    return (
    <div className="wrap">
      <table>
        <thead>
            <tr>
                <td>Categories</td>
                <td>Actions</td>
            </tr>
        </thead>
            <tbody>
                {rows}
            </tbody>
        </table>
        <button><Link to="/home/categories/add">Add Categories</Link></button>
    </div>
  );
    }
}
  

export default Categories;