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
    let start = '';
    let end ='';
    let clothes = '';
    let places = [];

    this.state.groups.forEach((item, i) => {
        places.push(item['placeName']);
        start = item['startTime'];
        end = item['endTime'];
        item['Clothes'].forEach((elem) => {
            clothes = `${clothes}${(clothes ? ', ' : '')}${elem['color']} ${elem['name']}`;
        });
        rows.push(<tr>
        <td>{start}</td>
        <td>{end}</td>
        <td>{places}</td>
        <td>{clothes}</td>
        <td><button className="green">insert</button><button className="red">del</button></td>
        </tr>);
        start = '';
        end = '';
        clothes = '';
        places = [];
    })

    return (
        <div className="wrap Day">
          <h3>График на день</h3>
            <table className="Day">
                <thead>
                    <tr>
                        <td>От</td>
                        <td>До</td>
                        <td>Место</td>
                        <td>Одежда</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>

                    {rows}
                    
                </tbody>
            </table>
        </div>
    );
  }
}
  

export default Day;