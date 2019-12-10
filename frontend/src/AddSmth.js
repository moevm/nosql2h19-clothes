import React, {Component} from 'react';
import './AddSmth.css';
import './Form.css';

class AddSmth extends Component {
  render() {
    console.log(this.props.match.path);
    let str = "";
    if (this.props.match.path === "/add_category") {
      str = "Category";
    } else {
      str = "Place";
    }
    return (
    <form>
      <h3>Форма добавления {str}</h3>
      {console.log(this.props.path)}
      <div className="form__container">
        <label>{str}</label><input type="text" placeholder={str}></input>
      </div>
      <button type="submit" className="addSmth_button">Add {str}</button>
    </form>
  );
    }
}
  

export default AddSmth;