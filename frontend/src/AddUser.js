import React, {Component} from 'react';
import './AddUser.css';

class AddUser extends Component {
  render() {
    return (
    <form>
      <h3>Форма добавления</h3>
      <input type="text" placeholder="Login"></input>
      <input type="text" placeholder="Password"></input>
      <button type="submit">Сохранить</button>
    </form>
  );
    }
}
  

export default AddUser;