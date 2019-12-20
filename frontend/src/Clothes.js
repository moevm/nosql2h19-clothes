import React, {Component} from 'react';
import axios from 'axios';
import {Link} from 'react-router-dom';
import './Table.css';
import './Clothes.css';

let endpoint = "http://localhost:5000";

class Clothes extends Component {
    constructor(props) {
        super(props);

        this.state = {
            clothes: []
        }
    }
    componentDidMount() {
        this.getClothes();
    }

    getClothes = () => {
        axios
            .get(endpoint + "/api/home/clothes")
            .then(res => {
                console.log(res);
                this.setState({
                    clothes: res.data.map(item => {
                        return item;
                    })
                })
            })
        //   fetch("http://localhost:5000/api/home/places")
        //       .then(data => console.log(data));
    }

    deleteClothe = (id) => {
        axios
            .delete(endpoint + `/api/home/clothes/${id}`)
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
        let category = '';
        let name ='';
        let style = '';
        this.state.clothes.forEach((item, i) => {
            category = item['categoryName'];
            name = item['name'];
            style = item['styleName'];
            // item['Clothes'].forEach((elem) => {
            //     clothes = `${clothes}${(clothes ? ', ' : '')}${elem['color']} ${elem['name']}`;
            // });
            rows.push(<tr key={i}>
                <td>{category}</td>
                <td>{name}</td>
                <td>{style}</td>
                <td><button className="green">insert</button><button className="red" onClick={(e) => this.deleteClothe(item['name'], e)}>del</button></td>
            </tr>);
            category = '';
            name = '';
            style = '';
        })
        return (
            <div className="wrap Clothes">
            <table>
            <thead>
            <tr>
            <td>Название</td>
            <td>Категория</td>
            <td>Стиль</td>
            <td>Actions</td>
            </tr>
            </thead>
            <tbody>
            {rows}
            </tbody>
            </table>
            <button><Link to="/home/places/add">Добавление одежды</Link></button>
        </div>
    );
    }
}


export default Clothes;