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

    ///データ読み込み(一旦usersテーブルを読み込んでいる)
    useEffect(() => {
        async function fetchData() {
            const response = await axios.get('users')
            setTodos(response.data)
            setIsLoaded(true)
        }

        fetchData();
    }, []);

    //編集機能(編集→再投稿でUI上で変化、変更するをClickで永続的に変化する)
    const changeTodo = async ( id:string,name:string ) => {
    await client.put(`users/${id}`,{name: `${name}` })
    client.get('users').then(({ data }) => {
        setTodos(data)
    })
    }

    const submitEdits = async (id:string) => {
        const updatedTodoLists = [...todos].map((todo:any) => {
            if (todo.ID === id) {
            todo.Name = editingText;
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


    return (
    <div className="p-4 border border-gray-200 rounded shadow-lg">
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
                <td className="p-1">{todo.ID}</td>
                <td className="p-1">{todo.Name}</td>
                {/* 永続的変更ボタン */}
                <td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => changeTodo( todo.ID,todo.Name )}
                    >
                    変更する
                    </button>
                </td>
                {/* 編集を押すと、三項演算子によって再投稿ボタンが出てくる */}
                {todo.ID === todoEditing ? (
                <td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => submitEdits( todo.ID )}
                    >
                    再投稿
                    </button>
                </td>
                ):(<td className="p-1">
                    <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => setTodoEditing( todo.ID )}
                    >
                    編集
                    </button>
                </td>
                )}
                {todoEditing === todo.ID ? (
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
                    onClick={() => deleteTodo( todo.ID )}
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

