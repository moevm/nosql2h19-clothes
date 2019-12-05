import React, {Component} from 'react';
import './ShowCat.css';
import {Link} from 'react-router-dom';

class ShowCat extends Component {
  render() {
    return (
    <div className="ShowCat">
        <div className="ShowCat__wrap">
            <div className="ShowCat__wrap__item">
                <img src="http://t0.gstatic.com/images?q=tbn%3AANd9GcSQJ_ofWAfU9rUkYw8F8rzNnH9psFiS-ZhIvJ0sUVV28xbi3vH_SkcNW3-bMA&usqp=CAc" alt="текст" />
                <p>Поло</p>
            </div>
            <div className="ShowCat__wrap__item">
                <img src="http://t0.gstatic.com/images?q=tbn%3AANd9GcSQJ_ofWAfU9rUkYw8F8rzNnH9psFiS-ZhIvJ0sUVV28xbi3vH_SkcNW3-bMA&usqp=CAc" alt="текст" />
                <p>Поло</p>
            </div>
            <div className="ShowCat__wrap__item">
                <img src="http://t0.gstatic.com/images?q=tbn%3AANd9GcSQJ_ofWAfU9rUkYw8F8rzNnH9psFiS-ZhIvJ0sUVV28xbi3vH_SkcNW3-bMA&usqp=CAc" alt="текст" />
                <p>Поло</p>
            </div>
        </div>
        <button className="ShowCat__btn"><Link to='/home/categories/show/add'>Добавить предмет</Link></button>
    </div>
  );
    }
}

  

export default ShowCat;