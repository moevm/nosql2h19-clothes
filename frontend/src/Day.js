import React, {Component} from 'react';
import './Day.css';

class Day extends Component {
  render() {
    return (
        <div className="wrap Day">
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
                    <tr>
                        <td>13:00 - 16:00</td>
                        <td>Работа</td>
                        <td>Кроссовки, джинсы, кофта</td>
                        <td><button className="green">insert</button><button className="red">del</button></td>
                    </tr>
                    <tr>
                        <td>16:00 - 18:00</td>
                        <td>Театр</td>
                        <td>Кроссовки, джинсы, кофта</td>
                        <td><button className="green">insert</button><button className="red">del</button></td>
                    </tr>
                    
                </tbody>
            </table>
        </div>
    );
  }
}
  

export default Day;