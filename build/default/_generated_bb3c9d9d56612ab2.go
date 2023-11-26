components {
  id: "player"
  component: "/main/player.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
components {
  id: "bullet_shot"
  component: "/main/bullet_shot.sound"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
components {
  id: "background_music"
  component: "/main/background_music.sound"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
components {
  id: "explosion_particle"
  component: "/main/explosion.particlefx"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
components {
  id: "explosion"
  component: "/main/explosion.sound"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/game.atlas\"\ndefault_animation: \"player_4\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 1.5
    y: 1.5
    z: 1.5
  }
}
embedded_components {
  id: "bullet"
  type: "factory"
  data: "prototype: \"/main/bullet.go\"\nload_dynamically: false\ndynamic_prototype: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"player\"\nmask: \"enemy_bullets\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 8.541181\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 6.415932\n      y: -4.27059\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 1\n    count: 1\n  }\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 14.11505\n      y: -12.811771\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 2\n    count: 1\n  }\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: -6.415932\n      y: -4.27059\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: -0.07361815\n      w: 0.9972865\n    }\n    index: 3\n    count: 1\n  }\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: -14.115051\n      y: -11.388241\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: -0.07361815\n      w: 0.9972865\n    }\n    index: 4\n    count: 1\n  }\n  data: 7.8688526\n  data: 7.8688526\n  data: 7.8688526\n  data: 7.8688526\n  data: 7.8688526\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\nbullet: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
