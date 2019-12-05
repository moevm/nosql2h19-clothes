import React, {Component} from 'react';
import './Make.css';

class Make extends Component {
  render() {
    return (
        <form className="Make">
            <h3>Генерация графика</h3>
            <div>Укажите предпочтения одежды для своих мест</div>
            <div>
              <select >
                <option>Place 1</option>
                <option>Place 2</option>
                <option>Place 3</option>
              </select>
              <select>
                <option>Style 1</option>
                <option>Style 2</option>
                <option>Style 3</option>
              </select>
            </div>
            <button type="submit">Generate</button>
        </form>
    );
  }
}
  

export default Make;