import React, {Component} from 'react';
import axios from 'axios';
import './Day.css';

let endpoint = "http://localhost:5000";

class Day extends Component {
    constructor(props) {
        super(props);

        this.state = {
            groups: []
        }
    }
    componentDidMount() {
        this.getGroups();
    }

    getGroups = () => {
        axios
            .get(endpoint + "/api/home/groups")
            .then(res => {
                console.log(res);
                this.setState({
                    groups: res.data.map(item => {
                        return item;
                    })
                })
            })
    }
  render() {
    let rows = [];
    let times = '';
    let clothes = '';
    let places = [];

    this.state.groups.forEach((item, i) => {
        places.push(item['place_name']);
        times = item['date'];
        item['Clothes'].forEach((elem) => {
            clothes = `${clothes}${(clothes ? ', ' : '')}${elem['color']} ${elem['name']} ${elem['notes']}`;
        });
        rows.push(<tr>
        <td>{times}</td>
        <td>{places}</td>
        <td>{clothes}</td>
        <td><button className="green">insert</button><button className="red">del</button></td>
        </tr>);
        times = '';
        clothes = '';
        places = [];
    })

    return (
        <div className="wrap Day">
          <h3>График на день</h3>
            <table className="Day">
                <thead>
                    <tr>
                        <td>Время</td>
                        <td>Место</td>
                        <td>Одежда</td>
                        <td>Actions</td>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>9:00 - 13:00</td>
                        <td>Институт</td>
                        <td>Кроссовки, джинсы, кофта</td>
                        <td><button className="green">insert</button><button className="red">del</button></td>
                    </tr>
                    {rows}
                    
                </tbody>
            </table>
        </div>
    );
  }
}
  

export default Day;