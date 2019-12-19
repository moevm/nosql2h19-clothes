import React, {Component} from 'react';
import './AddSmth.css';
import './Form.css';
import axios from 'axios';
import {Link} from "react-router-dom";

let endpoint = "http://localhost:5000";

class AddPlace extends Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleChange = this.handleChange.bind(this);
  }

  handleChange = (event) => {
    this.setState({value: event.target.value});
    // console.log(event.target.value);
    // console.log(this.state.value);
  }

  addPlace = () => {
    // console.log(this.state.value);
    axios
        .post(endpoint + "/api/home/places", {
          name: this.state.value
        })
        .then((response) => {
          alert(response.status);
        })
}
  render() {
    return (
    <div className="AddPlace">
      <h3>Форма добавления Places</h3>
      <div className="form__container">
        <label>Places</label><input type="text" placeholder="Place" id="place" onChange={this.handleChange}></input>
      </div>
    <div className="wrap-btn">
      <button className="addSmth_button"><Link to="/home/places/">Back</Link></button>
      <button type="submit" className="addSmth_button" onClick={this.addPlace}>Add Places</button>
        </div>
    </div>
  );
    }
}
  

export default AddPlace;