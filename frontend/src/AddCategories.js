import React, {Component} from 'react';
import './AddSmth.css';
import './Form.css';

class AddCategories extends Component {
  render() {
    return (
    <form>
      <h3>Форма добавления Categories</h3>
      <div className="form__container">
        <label>Category</label><input type="text" placeholder="Category"></input>
      </div>
      <button type="submit" className="addSmth_button">Add Category</button>
    </form>
  );
    }
}
  

export default AddCategories;