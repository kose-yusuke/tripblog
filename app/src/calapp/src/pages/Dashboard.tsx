import { InputForm } from '../components/InputForm'
import  TodoList  from '../components/Todolist'

const TodoDashboard = () => {
    return (
    <div className="w-128 mx-auto my-20">
        <InputForm />
        <TodoList />
    </div>
    )
}

export default TodoDashboard;