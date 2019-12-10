import React, {Component} from 'react';
import './EditPlace.css';

class EditPlace extends Component {
  render() {
    return (
    <form className="Make EditPlace">
      <h3>Настройка места посещения</h3>
        <div>Устанавливается и настраивается время и даты посещения мест</div>
        <div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Понедельник</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Вторник</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Среда</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Четверг</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Пятница</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Суббота</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
            <div className="EditPlace__wrap">
                <div className="EditPlace__wrap wrap__date">
                    <input type="checkbox" />
                    <p>Воскресенье</p>
                </div>
                <div className="EditPlace__wrap wrap__time">
                    <p>С</p>
                    <input type="time" />
                    <p>До</p>
                    <input type="time" />
                </div>
            </div>
        </div>
        <button type="submit">Save</button>
    </form>
  );
    }
}

  

export default EditPlace;