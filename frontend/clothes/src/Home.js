import React from 'react';
import './index.css';
import {Link} from 'react-router-dom';


class Home extends React.Component {
    constructor() {
        const date = new Date();
        let month = date.getMonth();
        let year = date.getFullYear();
        super()
        this.state = {
            month: month,
            year: year
        }
        this.handleClick = this.handleClick.bind(this)
    }

    handleClick() {
        this.setState(prevState => {
            return {
                month: prevState.month - 1
            }
        })
    }
    render() {
        const date = new Date();
        let month = date.getMonth();
        let year = date.getFullYear();
        let current = new Date(year, month, 1);
        let next = new Date(year, month + 1, 1);
        let diff = (next - current) / (1000 * 3600 * 24); 
        const ROWS = 6;
        const COLS = 7;
        let table = [];
        let tr = [];
        let index = (current.getDay() + 6) % 7;
        let k = 1 - index;
        const monthNames = [ "Январь","Февраль","Март","Апрель","Май","Июнь",
    "Июль","Август","Сентябрь","Октябрь","Ноябрь","Декабрь" ];

        for (let i = 0; i < ROWS; i++) {
            tr = [];
            for (let j = 0; j < COLS; j++) {
                tr.push(<td style={k === date.getDate() ? {backgroundColor: "#9ca6e4"} : {backgroundColor: "#ffffff"}}><Link to='/home/day'>{(k > 0 && k <= diff) ? k : ''}</Link></td>);
                k += 1;
            }
            table.push(<tr>{tr}</tr>)
        }

        const leftArrow = () => {
            month -= 1;
            console.log(month);
        }

    
        return (
            <table>
                <thead>
                    <td colSpan="7"><button onClick={this.handleClick}>&lt;</button>{monthNames[this.state.month]} {year}<button>&gt;</button></td>
                </thead>
                <tr>
                    <td>Пн</td>
                    <td>Вт</td>
                    <td>Ср</td>
                    <td>Чт</td>
                    <td>Пт</td>
                    <td>Сб</td>
                    <td>Вс</td>
                </tr>
                {table}
            </table>
        );
    }
}


export default Home;