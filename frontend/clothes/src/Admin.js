import React, {Component} from 'react';
import './Admin.css';
import {Link} from 'react-router-dom';

class Admin extends Component {
  render() {
    return (
        <div className="wrap Admin">
            <table className="Admin">
                <thead>
                    <tr>
                        <td>User</td>
                        <td>Actions</td>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>9:00 - 13:00</td>
                        <td><button className="green"><Link to='/admin/edit'>fix</Link></button><button className="red">del</button></td>
                    </tr>
                    <tr>
                        <td>13:00 - 16:00</td>
                        <td><button className="green"><Link to='/admin/edit'>fix</Link></button><button className="red">del</button></td>
                    </tr>
                    <tr>
                        <td>16:00 - 18:00</td>
                        <td><button className="green"><Link to='/admin/edit'>fix</Link></button><button className="red">del</button></td>
                    </tr>
                </tbody>
            </table>
            <div className="wrap-btn Admin">
                <button><Link to='/home/places'>Статистика</Link></button>
                <button><Link to='/home/categories'>Add user</Link></button>
            </div>
        </div>
    );
  }
}
  

export default Admin;