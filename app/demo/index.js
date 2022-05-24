
var page = 1
var pageSize = 20
window.onload = (h, e) => {
    init()
}

const baseURL = "http://124.221.160.98:4200"

const init = async () => {
    const sgnList = document.getElementById("ntfList");
    chidren = [...sgnList.childNodes]
    chidren.forEach(n => sgnList.removeChild(n))
    const response = await fetch(`${baseURL}/api/v1/sgn?page=${page}&page_size=${pageSize}`)
    const data = await response.json()
    if (data.success && data.data) {
        data.data.forEach(ntf => {
            sgnList.appendChild(createNtfItem(ntf))
        });
    }
}

const createNtfItem = (data) => {
    const item = document.createElement('div')
    item.className = 'sgn-item'
    const name = document.createElement('div')
    name.className = 'sgn-item-name'
    name.innerText = data.token_id
    const img = document.createElement('img')
    img.className = 'sgn-item-img'
    img.src = `${baseURL}/api/v1/blob/${encodeURIComponent(data.metadata.image)}?w=120&h=120`
    item.appendChild(name)
    item.appendChild(img)
    return item
}

const prePage = async () => {
    page--
    if (page < 1) {
        page = 1
    }
    await init()
}

const nextPage = async () => {
    page++
    await init()
}

const setPageSize = async (size) => {
    if (pageSize != size) {
        pageSize = size
        await init();
    }
}

const submitPageSize = async () => {
    const size = document.getElementById("pagesizeInput").value
    if (size) {
        await setPageSize(size)
    }
}