import React from 'react';
import './index.css';
import './Home.css';
import './Table.css';
import {Link} from 'react-router-dom';


class Home extends React.Component {
    constructor() {
        const date = new Date();
        let month = date.getMonth();
        let year = date.getFullYear();
    //     const date = new Date();
    //     let month = date.getMonth();
    //     let year = date.getFullYear();
        let current = new Date(year, month, 1);
        let next = new Date(year, month + 1, 1);
    // let diff = (next - current) / (1000 * 3600 * 24); 
    //     const ROWS = 6;
    //     const COLS = 7;
    //     let table = [];
    //     let tr = [];
    //     let index = (current.getDay() + 6) % 7;
    //     let k = 1 - index;
    //     const monthNames = [ "Январь","Февраль","Март","Апрель","Май","Июнь",
    // "Июль","Август","Сентябрь","Октябрь","Ноябрь","Декабрь" ];
        super()
        this.state = {
            date: date,
            month: month,
            year: year,
            current: current,
            next: next

        }
        this.handleClickLeft = this.handleClickLeft.bind(this);
        this.handleClickRight = this.handleClickRight.bind(this);
    }

    handleClickLeft() {
        // table = [];
        this.setState(prevState => {
            return {
                month: prevState.month - 1,
                current: new Date(prevState.year, prevState.month - 1, 1),
                next: new Date(prevState.year, prevState.month, 1)
            }
        })
    }

    handleClickRight() {
        // table = [];
        this.setState(prevState => {
            return {
                month: prevState.month + 1,
                current: new Date(prevState.year, prevState.month + 1, 1),
                next: new Date(prevState.year, prevState.month + 2, 1)
            }
        })
    }

    render() {
        // const date = new Date();
        // let month = date.getMonth();
        // let year = date.getFullYear();
        // let current = new Date(year, month, 1);
        // let next = new Date(year, month + 1, 1);
        let diff = (this.state.next - this.state.current) / (1000 * 3600 * 24); 
        const ROWS = 6;
        const COLS = 7;
        let table = [];
        let tr = [];
        let index = (this.state.current.getDay() + 6) % 7;
        let k = 1 - index;
        const monthNames = [ "Январь","Февраль","Март","Апрель","Май","Июнь",
    "Июль","Август","Сентябрь","Октябрь","Ноябрь","Декабрь" ];

    console.log(k);    

        for (let i = 0; i < ROWS; i++) {
            tr = [];
            for (let j = 0; j < COLS; j++) {
                tr.push(<td style={{backgroundColor: "#ffffff"}}><Link to='/home/day'>{(k > 0 && k <= diff) ? k : ''}</Link></td>);
                k += 1;
            }
            table.push(<tr>{tr}</tr>);
            console.log("ok");
        }

        console.log(k);

        // const leftArrow = () => {
        //     month -= 1;
        //     console.log(month);
        // }

    
        return (
            <div className="wrap">
                <table>
                    <thead>
                        <tr>
                            <td colSpan="7"><button onClick={this.handleClickLeft}>&lt;</button>{monthNames[this.state.month]} {this.state.year}<button onClick={this.handleClickRight}>&gt;</button></td>
                        </tr>
                    </thead>
                    <tbody>
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
                    </tbody>
                </table>
                <div className="wrap-btn">
                    <button><Link to='/home/places'>Мои места</Link></button>
                    <button><Link to='/home/categories'>Мои категории</Link></button>
                    <button><Link to='/home/make'>Генерация графика</Link></button>
                </div>
            </div>
        );
    }
}


export default Home;