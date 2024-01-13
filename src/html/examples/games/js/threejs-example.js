// set up the scene, camera, and renderer
const scene = new THREE.Scene()
const camera = new THREE.PerspectiveCamera(
  75,
  window.innerWidth / window.innerHeight,
  0.1,
  1000
)
const renderer = new THREE.WebGLRenderer()
renderer.setSize(window.innerWidth, window.innerHeight)
document.body.appendChild(renderer.domElement)

// create a cube and add it to the scene
const geometry = new THREE.BoxGeometry()
const material = new THREE.MeshBasicMaterial({
  color: 0x00ff00,
})
const cube = new THREE.Mesh(geometry, material)
scene.add(cube)

// set the initial position of the cube
cube.position.set(0, 0, -5)

// handle user input for cube movement
const moveSpeed = 0.1
document.addEventListener('keydown', (event) => {
  switch (event.key) {
    case 'ArrowUp':
      cube.position.z += moveSpeed
      break
    case 'ArrowDown':
      cube.position.z -= moveSpeed
      break
    case 'ArrowLeft':
      cube.position.x -= moveSpeed
      break
    case 'ArrowRight':
      cube.position.x += moveSpeed
      break
  }
})

// set up the camera position
camera.position.z = 5

// render loop
function animate() {
  requestAnimationFrame(animate)
  renderer.render(scene, camera)
}
animate()
