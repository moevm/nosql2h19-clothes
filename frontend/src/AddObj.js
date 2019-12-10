import React, {Component} from 'react';
import './AddObj.css';

class AddObj extends Component {
  render() {
    return (
    <form>
      <h3>Форма добавления</h3>
      <input type="text" placeholder="Загрузить картинку"></input>
      <textarea placeholder="Description"></textarea>
      <button type="submit">Сохранить</button>
    </form>
  );
    }
}
  

export default AddObj;