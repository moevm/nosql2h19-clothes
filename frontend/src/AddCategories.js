import React, {Component} from 'react';
import './AddSmth.css';
import './Form.css';
import axios from 'axios';
import {Link} from "react-router-dom";

let endpoint = "http://localhost:5000";

class AddCategories extends Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleChange = this.handleChange.bind(this);
  }

  handleChange = (event) => {
    this.setState({value: event.target.value});
  }

  addCategory = () => {
    axios
        .post(endpoint + "/api/home/categories", {
          name: this.state.value
        })
        .then((response) => {
          alert(response.status);
        })
  }

  render() {
    return (
    <div className="AddPlace">
      <h3>Форма добавления Категорий</h3>
      <div className="form__container">
        <label>Category</label><input type="text" placeholder="Категория" onChange={this.handleChange}></input>
      </div>
      <div className="wrap-btn">
        <button className="addSmth_button"><Link to="/home/categories/">Back</Link></button>
        <button type="submit" className="addSmth_button" onClick={this.addCategory}>Добавить категорию</button>
      </div>
    </div>
  );
    }
}
  

export default AddCategories;