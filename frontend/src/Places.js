import React, {Component} from 'react';
import axios from 'axios';
import {Link} from 'react-router-dom';
import './Table.css';

let endpoint = "http://localhost:5000";

class Places extends Component {
  constructor(props) {
      super(props);

      this.state = {
          places: []
      }
  }
  componentDidMount() {
      this.getPlaces();
  }
    
  getPlaces = () => {
    axios
      .get(endpoint + "/api/home/places")
      .then(res => {
          console.log(res);
          this.setState({
              places: res.data.map(item => {
                  return item['name'];
              })
          })
      })
    //   fetch("http://localhost:5000/api/home/places")
    //       .then(data => console.log(data));
  }

  deletePlace = (id) => {
      axios
          .delete(endpoint + `/api/home/places/${id}`)
          .then(res => {
              console.log(res);
              console.log(res.data);
              document.location.reload(true);
          })
  }

  pageReload = () => {
      document.location.reload(true);
  }

  render() {
    let rows = [];
    this.state.places.forEach((item, i) => {
        rows.push(<tr key={i}><td>{item}</td>
        <td><button className="green" onClick={this.pageReload}>see</button><button className="red" onClick={(e) => this.deletePlace(item, e)}>del</button></td></tr>)
    })
    return (
    <div className="wrap">
      <table>
        <thead>
            <tr>
                <td>Places</td>
                <td>Actions</td>
            </tr>
        </thead>
        <tbody>
          {rows}
          </tbody>
        </table>
        <button><Link to="/home/places/add">Add Place</Link></button>
    </div>
  );
    }
}
  

export default Places;