import React, {Component} from 'react';
import './Admin.css';
import {Link} from 'react-router-dom';
import axios from 'axios';

let endpoint = "http://localhost:5000";

class Admin extends Component {

    constructor(props) {
        super(props);

        this.state = {
            users: [],
            filename: ''
        }
    }

    componentDidMount() {
        this.getUsers();
    }

    getUsers = () => {
        axios
            .get(endpoint + "/api/admin/users")
            .then(res => {
                console.log(res.data);
                this.setState({
                    users: res.data.map(item => {
                        return item;
                    })
                })
            })
    };

    deleteUser = (id) => {
        console.log(id)
        axios

            .delete(endpoint + `/api/admin/${id}`)
            .then(res => {
                console.log(res);
                console.log(res.data);
                document.location.reload(true);
            })
    };

    getFileName = () => {
        axios
            .get(endpoint + "/api/admin")
            .then(res => {
                console.log(res.data);
                this.setState({
                    filename: res.data
                })
            })
    }

    sendFile = () => {
        let formData = new FormData();
        formData.append("myFile", document.getElementById("file").files[0]);

        axios
            .post(endpoint + "/api/admin", formData)
            .then((response) => {
                alert(response.status);
            })
    }

    downloadFile = () => {
        let filename_ = this.getFileName();
        let filename = 'mydb.json'
        let link = endpoint + '/api/tmp/';
        const dummy = document.createElement("a");
        dummy.href = link + filename;
        dummy.download = filename;

        document.body.appendChild(dummy);
        dummy.click();
    };
  render() {
      let rows = [];
      let email = '';
      let name ='';
      let age = '';

      this.state.users.forEach((item, i) => {
          // places.push(item['placeName']);
          email = item['Email'];
          name = item['Name'];
          age = item['Age'];
          // item['Clothes'].forEach((elem) => {
          //     clothes = `${clothes}${(clothes ? ', ' : '')}${elem['color']} ${elem['name']}`;
          // });
          rows.push(<tr key={i}>
          <td>{email}</td>
          <td>{name}</td>
          <td>{age}</td>
          <td><button className="green">see</button><button className="red" onClick={(e) => this.deleteUser(item['Email'])}>del</button></td>
          </tr>);
          email = '';
          name = '';
          age = '';
      })
    return (
        <div className="wrap Admin">
            <table className="Admin">
                <thead>
                    <tr>
                        <td>Email</td>
                        <td>Name</td>
                        <td>Age</td>
                        <td>Actions</td>
                    </tr>
                </thead>
                <tbody>
                    {rows}
                </tbody>
            </table>
            <div className="wrap-btn Admin">
                <button><Link to='/home/places'>Статистика</Link></button>
                <button><Link to='/home/categories'>Add user</Link></button>
               <button onClick={this.downloadFile}>Скачать файл</button>
                <input type="file" id="file"></input>
                <button onClick={this.sendFile}>Импорт</button>
            </div>
        </div>
    );
  }
}
  

export default Admin;