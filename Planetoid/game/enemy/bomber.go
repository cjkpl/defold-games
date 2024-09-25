components {
  id: "script"
  component: "/game/enemy/bomber.script"
}
components {
  id: "exp"
  component: "/game/effects/exp-bomber.particlefx"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bomber\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game/core/game.tilesource\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"player\"\n"
  "mask: \"laser\"\n"
  "mask: \"smart\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 3.0\n"
  "  data: 3.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "bomb-factory"
  type: "factory"
  data: "prototype: \"/game/enemy/bomb.go\"\n"
  ""
}
