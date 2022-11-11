import axios from 'axios'
import { useEffect, useState } from 'react'
import { client } from '../libs/axios'

const TodoList = (props:any) => {
    //Todo状態管理
    const [todos, setTodos ] = useState(Object)
    const [isLoaded, setIsLoaded] = useState(false)
	const [error, setError] = useState("")
    //編集状態管理
    const [todoEditing, setTodoEditing] = useState(null);
    const [editingText, setEditingText] = useState("");
    const [todoName, setTodoName] = useState('')

    ///データ読み込み(一旦usersテーブルを読み込んでいる)
    useEffect(() => {
        async function fetchData() {
            const response = await axios.get('users')
            setTodos(response.data)
            setIsLoaded(true)
        }

        fetchData();
    }, []);

    //追加機能
    const addTodo = async () => {
        const body = {
            id : String(Math.floor( Math.random() * 1000 ) + 1),
            name: todoName
        }
        console.log(body)
        await client.post('users', body)
        client.get('users').then(({ data }) => {
            setTodos(data)
        })
        }

    //編集機能(編集→再投稿でUI上で変化、変更するをClickで永続的に変化する)
    const changeTodo = async ( id:string,name:string ) => {
    await client.put(`users/${id}`,{name: `${name}` })
    client.get('users').then(({ data }) => {
        setTodos(data)
    })
    }

    const submitEdits = async (id:string) => {
        const updatedTodoLists = [...todos].map((todo:any) => {
            if (todo.id === id) {
            todo.name = editingText;
            }
            return todo;
        });
        setTodos(updatedTodoLists);
        setTodoEditing(null);
        setEditingText("");
    };

    //idを元にdelete
    const deleteTodo = async (id:string) => {        
        await client.delete(`users/${id}`)
        client.get('users').then(({ data }) => {
        setTodos(data)
    })
    }

    //Loading,UX向上
    if (!isLoaded) {
		return (
			<p>Loading...</p>
		)
	}

	if (error) {	
		return (
			<div>Error: {error}</div>
			)
	} else if (!isLoaded) {
		return (
			<p>Loading...</p>
		)
	}

    //API出力確認用
    console.log(todos)
    //inputフォーム
    const onChangeTodoName: React.ChangeEventHandler<HTMLInputElement> = (
        event
        ) => {
        setTodoName(event.target.value)
        }




    return (
    <div className="p-4 border border-gray-200 rounded shadow-lg">
        <p className="mb-2 font-bold">新しいタスクを追加する</p>
        <input
        placeholder="買い物"
        className="mr-4 border shadow-md border-teal-500 rounded"
        onChange={onChangeTodoName}
        />
        <button
        onClick={() => addTodo() }
        className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
        >
        追加
        </button>
        <p className="font-bold mb-2">タスク一覧</p>
        <table className="border-collapse table-auto">
        <thead>
            <tr>
                <th className="py-1">番号</th>
                <th className="py-1">ID</th>
                <th className="p-1">タスク名</th>
            </tr>
        </thead>
        <tbody>
            {todos?.map((todo:any, index:any) => {
            return (
                <tr key={index}>
                <td className="p-1">{index}</td>
                <td className="p-1">{todo.id}</td>
                <td className="p-1">{todo.name}</td>
                {/* 永続的変更ボタン */}
                <td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => changeTodo( todo.id,todo.name )}
                    >
                    変更する
                    </button>
                </td>
                {/* 編集を押すと、三項演算子によって再投稿ボタンが出てくる */}
                {todo.id === todoEditing ? (
                <td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => submitEdits( todo.id )}
                    >
                    再投稿
                    </button>
                </td>
                ):(<td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => setTodoEditing( todo.id )}
                    >
                    編集
                    </button>
                </td>
                )}
                {todoEditing === todo.id ? (
                    <input
                    type="text"
                    placeholder="編集内容を入力"
                    className="m-7 p-3 w-4/5 border-2"
                    value={editingText}
                    onChange={(e) => setEditingText(e.target.value)}
                    />
                ) : (
                    <div>{""}</div>
                )}
                {/* 削除 */}
                <td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => deleteTodo( todo.id )}
                    >
                    削除する
                    </button>
                </td>
                </tr>
            )
            })}
        </tbody>
        </table>

    </div>
    )
}

export default TodoList;

