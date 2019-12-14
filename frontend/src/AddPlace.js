import React, {Component} from 'react';
import './AddSmth.css';
import './Form.css';

class AddPlace extends Component {
  render() {
    return (
    <form>
      <h3>Форма добавления Places</h3>
      <div className="form__container">
        <label>Places</label><input type="text" placeholder="Places"></input>
      </div>
      <button type="submit" className="addSmth_button">Add Places</button>
    </form>
  );
    }
}
  

export default AddPlace;