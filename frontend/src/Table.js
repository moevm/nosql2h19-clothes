import React, {Component} from 'react';
import './Table.css';

class Table extends Component {
  render() {
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
                <tr>
                    <td>Институт</td>
                    <td><button className="green">see</button><button className="red">del</button></td>
                </tr>
                <tr>
                    <td>Работа</td>
                    <td><button className="green">see</button><button className="red">del</button></td>
                </tr>
                <tr>
                    <td>Театр</td>
                    <td><button className="green">see</button><button className="red">del</button></td>
                </tr>
            </tbody>
        </table>
        <button>Add Place</button>
    </div>
  );
    }
}
  

export default Table;