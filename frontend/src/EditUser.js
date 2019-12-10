import React, {Component} from 'react';
import './EditUser.css';

class EditUser extends Component {
  render() {
    return (
    <form>
      <h3>Форма редактирования</h3>
      <input type="text" placeholder="Login"></input>
      <input type="text" placeholder="Password"></input>
      <div className="wrap-btn Edit">
                <button>Удалить</button>
                <button>Обновить</button>
      </div>
    </form>
  );
    }
}
  

export default EditUser;