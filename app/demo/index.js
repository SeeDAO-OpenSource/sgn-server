
window.onload = (h, e) => {
    init()
}

const baseURL="http://localhost:5000"

const init = async () => {
    const response = await fetch(baseURL+'/api/v1/nft/0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1?page=1&pageSize=10')
    const data = await response.json()
    const ntfList = document.getElementById("ntfList");
    if (data.success) {
        data.data.forEach(ntf => {
            ntfList.appendChild(createNtfItem(ntf))
        });
    }
}

const createNtfItem = (data) => {
    const item = document.createElement('div')
    item.className = 'nft-item'
    const name = document.createElement('div')
    name.className = 'nft-item-name'
    name.innerText = data.token_id
    const img = document.createElement('img')
    img.className = 'nft-item-img'
    img.src = `${baseURL}/api/v1/blob/${encodeURIComponent(data.metadata.image)}?w=120&h=120`
    item.appendChild(name)
    item.appendChild(img)
    return item
}