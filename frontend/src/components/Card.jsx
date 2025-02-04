import React from "react";

const Card = (props) => {
    const listToDo = props.listToDo.map((_todo)=>{
        return (
            <div className="todo-list">
                <div className="form-cBox">
                    <input className="cBox" type="checkbox" checked={_todo.completed} onChange={e => onChangeComplete(_todo.id, e.target.checked)}></input>
                    <span>{_todo.title}</span>
                </div>
                
                <button onClick={()=>removeHandle(_todo.id)}>
                    <svg xmlns="http://www.w3.org/2000/svg"   viewBox="0 0 24 24" fill="none" stroke="#cacde8" stroke-width="1" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                </button>
            </div>
        );
    });

    return (
        <>
            {listToDo}
        </>
    );
}


export default Card;